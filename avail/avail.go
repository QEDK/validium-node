package avail

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/avail/internal/config"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"golang.org/x/crypto/sha3"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"

	availTypes "github.com/0xPolygonHermez/zkevm-node/avail/types"
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

type HeaderRPCResponse struct {
	Result Header `json:"result"`
}

type DataProofRPCResponse struct {
	Result DataProof `json:"result"`
}

type DataProof struct {
	Root           string   `json:"root"`
	Proof          []string `json:"proof"`
	NumberOfLeaves uint     `json:"number_of_leaves"`
	LeafIndex      uint     `json:"leaf_index"`
	Leaf           string   `json:"leaf"`
}

type Header struct {
	ParentHash     types.Hash        `json:"parentHash"`
	Number         types.BlockNumber `json:"number"`
	StateRoot      types.Hash        `json:"stateRoot"`
	ExtrinsicsRoot types.Hash        `json:"extrinsicsRoot"`
	Digest         types.Digest      `json:"digest"`
	Extension      struct {
		V1 struct {
			Commitment struct {
				Rows       uint       `json:"rows"`
				Cols       uint       `json:"cols"`
				DataRoot   types.Hash `json:"dataRoot"`
				Commitment []types.U8 `json:"commitment"`
			} `json:"commitment"`
			AppLookup struct {
				Size  uint `json:"size"`
				Index []struct {
					AppId uint `json:"app_id"`
					Start uint `json:"start"`
				} `json:"index"`
			} `json:"app_lookup"`
		} `json:"V1"`
	} `json:"extension"`
}

type HeaderAsScale struct {
	ParentHash     types.Hash        `json:"parentHash"`
	Number         types.BlockNumber `json:"number"`
	StateRoot      types.Hash        `json:"stateRoot"`
	ExtrinsicsRoot types.Hash        `json:"extrinsicsRoot"`
	Digest         types.Digest      `json:"digest"`
	Extension      struct {
		V1 struct {
			Commitment struct {
				Rows       types.UCompact `json:"rows"`
				Cols       types.UCompact `json:"cols"`
				DataRoot   types.Hash     `json:"dataRoot"`
				Commitment []types.U8     `json:"commitment"`
			} `json:"commitment"`
			AppLookup struct {
				Size  types.UCompact `json:"size"`
				Index []struct {
					AppId types.UCompact `json:"app_id"`
					Start types.UCompact `json:"start"`
				} `json:"index"`
			} `json:"app_lookup"`
		} `json:"V1"`
	} `json:"extension"`
}

type Index []struct {
	AppId types.UCompact `json:"app_id"`
	Start types.UCompact `json:"start"`
}

type Digest struct {
	Logs []string `json:"logs"`
}

type DigestAsScale struct {
	Logs []types.BytesBare `json:"logs"`
}

// type RPCResponse struct {
// 	Result Header `json:"result"`
// 	Error  string `json:"error"`
// }

// The following example shows how submit data blob and track transaction status
func PostData(txData []byte) (*availTypes.BatchDAData, error) {
	var config config.Config

	err := config.GetConfig("/app/avail-config.json")
	if err != nil {
		return nil, fmt.Errorf("cannot get config:%w", err)
	}

	api, err := gsrpc.NewSubstrateAPI(config.ApiURL)
	if err != nil {
		return nil, fmt.Errorf("cannot get api:%w", err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, fmt.Errorf("cannot get metadata:%w", err)
	}

	log.Infof("⚡️ Prepared data for Avail:%d bytes", len(txData))
	appID := 0

	// if app id is greater than 0 then it must be created before submitting data
	if config.AppID != 0 {
		appID = config.AppID
	}

	newCall, err := types.NewCall(meta, "DataAvailability.submit_data", types.NewBytes(txData))
	if err != nil {
		return nil, fmt.Errorf("cannot create new call:%w", err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(newCall)

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, fmt.Errorf("cannot get block hash:%w", err)
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, fmt.Errorf("cannot get runtime version:%w", err)
	}

	keyringPair, err := signature.KeyringPairFromSecret(config.Seed, 42)
	if err != nil {
		return nil, fmt.Errorf("cannot create keypair:%w", err)
	}

	key, err := types.CreateStorageKey(meta, "System", "Account", keyringPair.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create storage key:%w", err)
	}

	var accountInfo types.AccountInfo
	ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		return nil, fmt.Errorf("cannot get latest storage:%w", err)
	}

	pendingExt, err := api.RPC.Author.PendingExtrinsics()
	if err != nil {
		return nil, fmt.Errorf("cannot get pending extrinsics:%w", err)
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
		return nil, fmt.Errorf("cannot sign:%w", err)
	}

	// Send the extrinsic
	sub, err := api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return nil, fmt.Errorf("cannot submit extrinsic:%w", err)
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
				return nil, fmt.Errorf("❌ Extrinsic dropped")
			} else if status.IsUsurped {
				return nil, fmt.Errorf("❌ Extrinsic usurped")
			} else if status.IsInvalid {
				return nil, fmt.Errorf("❌ Extrinsic invalid")
			}
		case <-timeout:
			return nil, fmt.Errorf("⌛️ Timeout of 100 seconds reached without getting finalized status for extrinsic")
		}
	}

	log.Infof("✅ Data submitted by sequencer:%d bytes against AppID %v sent with hash %#x", len(txData), appID, blockHash)

	var dataProof DataProof
	var batchHash [32]byte
	maxTxIndex := 1
	h := sha3.NewLegacyKeccak256()
	h.Write(txData)
	h.Sum(batchHash[:0])

	for i := 0; i < maxTxIndex; i++ {
		resp, err := http.Post("https://kate.avail.tools/rpc", "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"kate_queryDataProof\",\"params\":[%d, \"%#x\"]}", i, blockHash)))
		if err != nil {
			return nil, fmt.Errorf("cannot post header request:%v", err)
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("cannot read body:%v", err)
		}

		var dataProofResp DataProofRPCResponse
		json.Unmarshal(data, &dataProofResp)

		log.Infof("💿 received data proof:%+v", dataProofResp.Result)

		if dataProofResp.Result.Leaf == fmt.Sprintf("%#x", batchHash) {
			dataProof = dataProofResp.Result
			break
		}

		maxTxIndex = int(dataProofResp.Result.NumberOfLeaves)
	}

	log.Infof("💿 received data proof:%+v", dataProof)
	var batchDAData availTypes.BatchDAData
	batchDAData.Proof = dataProof.Proof
	batchDAData.Width = dataProof.NumberOfLeaves
	batchDAData.LeafIndex = dataProof.LeafIndex

	resp, err := http.Post("https://kate.avail.tools/rpc", "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"chain_getHeader\",\"params\":[\"%#x\"]}", blockHash)))
	if err != nil {
		return nil, fmt.Errorf("cannot post header request:%v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	log.Infof("👍 received header:%v", data)

	if err != nil {
		return nil, fmt.Errorf("cannot read body:%v", err)
	}

	var headerResp HeaderRPCResponse
	json.Unmarshal(data, &headerResp)

	batchDAData.BlockNumber = uint(headerResp.Result.Number)
	log.Infof("🟢 prepared DA data:%+v", batchDAData)

	return &batchDAData, nil

	// log.Infof("received header:%+v", headerResp.Result)
	// encodedHeader, err := encodeHeaderAsScale(headerResp.Result)
	// if err != nil {
	// 	return fmt.Errorf("cannot encode header:%w", err)
	// }
	// log.Infof("received header:%+v", encodedHeader)

	// header, err := api.RPC.Chain.GetHeader(blockHash)
	// log.Infof("received header:%+v", header)

	// if err != nil {
	// 	return fmt.Errorf("cannot get header:%+v", err)
	// }

	// destAddress, err := types.NewHashFromHexString(config.DestinationAddress)
	// if err != nil {
	// 	return fmt.Errorf("cannot decode destination address:%w", err)
	// }

	// log.Infof("destination domain: %v, destination address: %v", types.NewUCompactFromUInt(uint64(config.DestinationDomain)), destAddress)

	// dispatchDataRootCall, err := types.NewCall(meta, "NomadDABridge.try_dispatch_data_root", types.NewUCompactFromUInt(uint64(config.DestinationDomain)), destAddress, encodedHeader)

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
}

func encodeHeaderAsScale(h Header) (HeaderAsScale, error) {
	var headerAsScale HeaderAsScale

	headerAsScale.ParentHash = h.ParentHash
	headerAsScale.Number = h.Number
	headerAsScale.StateRoot = h.StateRoot
	headerAsScale.ExtrinsicsRoot = h.ExtrinsicsRoot
	headerAsScale.Digest = h.Digest
	// headerAsScale.Digest.Logs = make([]types.BytesBare, len(h.Digest.Logs))
	// for i, item := range h.Digest.Logs {
	// 	var err error
	// 	headerAsScale.Digest.Logs[i], err = hex.DecodeString(item[2:])
	// 	if err != nil {
	// 		return headerAsScale, fmt.Errorf("cannot decode digest logs:%w", err)
	// 	}
	// }

	headerAsScale.Extension.V1.Commitment.Rows = types.NewUCompactFromUInt(uint64(h.Extension.V1.Commitment.Rows))
	headerAsScale.Extension.V1.Commitment.Cols = types.NewUCompactFromUInt(uint64(h.Extension.V1.Commitment.Cols))
	headerAsScale.Extension.V1.Commitment.DataRoot = h.Extension.V1.Commitment.DataRoot
	headerAsScale.Extension.V1.Commitment.Commitment = h.Extension.V1.Commitment.Commitment

	headerAsScale.Extension.V1.AppLookup.Size = types.NewUCompactFromUInt(uint64(h.Extension.V1.AppLookup.Size))
	headerAsScale.Extension.V1.AppLookup.Index = make(Index, len(h.Extension.V1.AppLookup.Index))

	for i, item := range h.Extension.V1.AppLookup.Index {
		headerAsScale.Extension.V1.AppLookup.Index[i].AppId = types.NewUCompactFromUInt(uint64(item.AppId))
		headerAsScale.Extension.V1.AppLookup.Index[i].Start = types.NewUCompactFromUInt(uint64(item.Start))
	}

	return headerAsScale, nil
}
