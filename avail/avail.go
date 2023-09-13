package avail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/avail/internal/config"
	"github.com/0xPolygonHermez/zkevm-node/log"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// type DataLookupIndexItem struct {
// 	AppId big.Int `json:"app_id"`
// 	Start big.Int `json:"start"`
// }
// type DataLookup struct {
// 	Size  big.Int               `json:"size"`
// 	Index []DataLookupIndexItem `json:"index"`
// }

// type KateCommitment struct {
// 	Rows       big.Int    `json:"rows"`
// 	Cols       big.Int    `json:"cols"`
// 	DataRoot   types.Hash `json:"dataRoot"`
// 	Commitment []types.U8 `json:"commitment"`
// }

// type V1HeaderExtension struct {
// 	Commitment KateCommitment `json:"commitment"`
// 	AppLookup  DataLookup     `json:"app_lookup"`
// }

// type HeaderExtension struct {
// 	V1 V1HeaderExtension `json:"V1"`
// }

// type Header struct {
// 	ParentHash     types.Hash        `json:"parentHash"`
// 	Number         types.BlockNumber `json:"number"`
// 	StateRoot      types.Hash        `json:"stateRoot"`
// 	ExtrinsicsRoot types.Hash        `json:"extrinsicsRoot"`
// 	Digest         types.Digest      `json:"digest"`
// 	Extension      HeaderExtension   `json:"extension"`
// }

type RPCResponse struct {
	Result struct {
		ParentHash     types.Hash        `json:"parentHash"`
		Number         types.BlockNumber `json:"number"`
		StateRoot      types.Hash        `json:"stateRoot"`
		ExtrinsicsRoot types.Hash        `json:"extrinsicsRoot"`
		Digest         types.Digest      `json:"digest"`
		Extension      struct {
			V1 struct {
				Commitment struct {
					Rows       big.Int    `json:"rows"`
					Cols       big.Int    `json:"cols"`
					DataRoot   types.Hash `json:"dataRoot"`
					Commitment []types.U8 `json:"commitment"`
				} `json:"commitment"`
				AppLookup struct {
					Size  big.Int `json:"size"`
					Index []struct {
						AppId big.Int `json:"app_id"`
						Start big.Int `json:"start"`
					}
				} `json:"app_lookup"`
			} `json:"V1"`
		} `json:"extension"`
	} `json:"result"`
}

// type RPCResponse struct {
// 	Result Header `json:"result"`
// 	Error  string `json:"error"`
// }

// The following example shows how submit data blob and track transaction status
func PostData(txData []byte) error {
	var config config.Config

	err := config.GetConfig("/app/avail-config.json")
	if err != nil {
		return fmt.Errorf("cannot get config:%w", err)
	}

	api, err := gsrpc.NewSubstrateAPI(config.ApiURL)
	if err != nil {
		return fmt.Errorf("cannot get api:%w", err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return fmt.Errorf("cannot get metadata:%w", err)
	}

	log.Infof("⚡️ Prepared data for Avail:%d bytes", len(txData))
	appID := 0

	// if app id is greater than 0 then it must be created before submitting data
	if config.AppID != 0 {
		appID = config.AppID
	}

	newCall, err := types.NewCall(meta, "DataAvailability.submit_data", types.NewBytes(txData))
	if err != nil {
		return fmt.Errorf("cannot create new call:%w", err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(newCall)

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return fmt.Errorf("cannot get block hash:%w", err)
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return fmt.Errorf("cannot get runtime version:%w", err)
	}

	keyringPair, err := signature.KeyringPairFromSecret(config.Seed, 42)
	if err != nil {
		return fmt.Errorf("cannot create keypair:%w", err)
	}

	key, err := types.CreateStorageKey(meta, "System", "Account", keyringPair.PublicKey)
	if err != nil {
		return fmt.Errorf("cannot create storage key:%w", err)
	}

	var accountInfo types.AccountInfo
	ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		return fmt.Errorf("cannot get latest storage:%w", err)
	}

	pendingExt, err := api.RPC.Author.PendingExtrinsics()
	if err != nil {
		return fmt.Errorf("cannot get pending extrinsics:%w", err)
	}

	nonce := uint32(accountInfo.Nonce) + uint32(len(pendingExt))
	options := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(500),
		AppID:              types.NewUCompactFromUInt(uint64(appID)),
		TransactionVersion: rv.TransactionVersion,
	}

	err = ext.Sign(keyringPair, options)
	if err != nil {
		return fmt.Errorf("cannot sign:%w", err)
	}

	// Send the extrinsic
	sub, err := api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return fmt.Errorf("cannot submit extrinsic:%w", err)
	}

	defer sub.Unsubscribe()
	timeout := time.After(100 * time.Second)
	var blockHash types.Hash
out:
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockHash = status.AsInBlock
				break out
			} else if status.IsFinalized {
				blockHash = status.AsFinalized
				break out
			} else if status.IsDropped {
				return fmt.Errorf("❌ Extrinsic dropped")
			}
		case <-timeout:
			return fmt.Errorf("⌛️ Timeout of 100 seconds reached without getting finalized status for extrinsic")
		}
	}

	log.Infof("✅ Data submitted by sequencer:%d bytes against AppID %v sent with hash %#x", len(txData), appID, blockHash)

	resp, err := http.Post("https://kate.avail.tools/rpc", "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"chain_getHeader\",\"params\":[\"%#x\"]}", blockHash)))
	if err != nil {
		return fmt.Errorf("cannot post header request:%v", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	log.Infof("received header:%v", data)

	if err != nil {
		return fmt.Errorf("cannot read body:%v", err)
	}

	var headerResp RPCResponse
	json.Unmarshal(data, &headerResp)

	log.Infof("received header:%+v", headerResp.Result)

	// header, err := api.RPC.Chain.GetHeader(blockHash)
	// log.Infof("received header:%+v", header)

	if err != nil {
		return fmt.Errorf("cannot get header:%+v", err)
	}

	// dispatchDataRootCall, err := types.NewCall(meta, "NomadDABridge.try_dispatch_data_root", types.NewU32(config.DestinationDomain), config.DestinationAddress, headerResp.Result)

	// if err != nil {
	// 	return fmt.Errorf("cannot create new call:%w", err)
	// }

	// dispatchDataRootExt := types.NewExtrinsic(dispatchDataRootCall)

	// nonce++
	// options = types.SignatureOptions{
	// 	BlockHash:          genesisHash,
	// 	Era:                types.ExtrinsicEra{IsMortalEra: false},
	// 	GenesisHash:        genesisHash,
	// 	Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
	// 	SpecVersion:        rv.SpecVersion,
	// 	Tip:                types.NewUCompactFromUInt(100),
	// 	AppID:              types.NewUCompactFromUInt(uint64(appID)),
	// 	TransactionVersion: rv.TransactionVersion,
	// }
	// err = dispatchDataRootExt.Sign(keyringPair, options)
	// if err != nil {
	// 	return fmt.Errorf("cannot sign:%w", err)
	// }

	// dispatchDataRootHash, err := api.RPC.Author.SubmitAndWatchExtrinsic(dispatchDataRootExt)
	// if err != nil {
	// 	return fmt.Errorf("cannot dispatch data root:%w", err)
	// }

	// log.Infof("✅ Data root dispatched by sequencer with AppID %v sent with hash %#x\n", appID, dispatchDataRootHash)

	return nil
}
