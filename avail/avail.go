// SPDX-License-Identifier: Apache-2.0
package avail

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	availConfig "github.com/0xPolygonHermez/zkevm-node/avail/internal/config"
	availTypes "github.com/0xPolygonHermez/zkevm-node/avail/types"
	"github.com/0xPolygonHermez/zkevm-node/log"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"golang.org/x/crypto/sha3"
)

// AccountNextIndexRPCResponse represents the next index of account rpc response
type AccountNextIndexRPCResponse struct {
	Result uint `json:"result"`
}

// DataProofRPCResponse represents the data proof rpc response
type DataProofRPCResponse struct {
	Result DataProof `json:"result"`
}

// DataProof represents the data proof structure
type DataProof struct {
	Root           string   `json:"root"`
	Proof          []string `json:"proof"`
	NumberOfLeaves uint     `json:"numberOfLeaves"`
	LeafIndex      uint     `json:"leafIndex"`
	Leaf           string   `json:"leaf"`
}

// nolint : revive
var (
	Config             availConfig.Config
	Api                *gsrpc.SubstrateAPI
	Meta               *types.Metadata
	AppId              int
	GenesisHash        types.Hash
	Rv                 *types.RuntimeVersion
	KeyringPair        signature.KeyringPair
	DestinationAddress types.Hash
	DestinationDomain  types.UCompact
)

const (
	networkID  = 42
	defaultTip = 1000
)

func init() {
	err := Config.GetConfig("/app/avail-config.json")
	if err != nil {
		log.Fatalf("cannot get config:%w", err)
	}

	Api, err = gsrpc.NewSubstrateAPI(Config.ApiURL)
	if err != nil {
		log.Fatalf("cannot get api:%w", err)
	}

	Meta, err = Api.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Fatalf("cannot get metadata:%w", err)
	}

	AppId = 0

	// if app id is greater than 0 then it must be created before submitting data
	if Config.AppID != 0 {
		AppId = Config.AppID
	}

	GenesisHash, err = Api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		log.Fatalf("cannot get block hash:%w", err)
	}

	Rv, err = Api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		log.Fatalf("cannot get runtime version:%w", err)
	}

	KeyringPair, err = signature.KeyringPairFromSecret(Config.Seed, networkID)
	if err != nil {
		log.Fatalf("cannot create keypair:%w", err)
	}

	DestinationAddress, err = types.NewHashFromHexString(Config.DestinationAddress)
	if err != nil {
		log.Fatalf("cannot decode destination address:%w", err)
	}

	DestinationDomain = types.NewUCompactFromUInt(uint64(Config.DestinationDomain))
}

// PostData posts data to the avail DA
func PostData(txData []byte) (*availTypes.BatchDAData, error) {
	log.Infof("⚡️ Prepared data for Avail:%d bytes", len(txData))

	newCall, err := types.NewCall(Meta, "DataAvailability.submit_data", types.NewBytes(txData))
	if err != nil {
		return nil, fmt.Errorf("cannot create new call:%w", err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(newCall)

	nonce, err := GetAccountNextIndex()
	if err != nil {
		return nil, fmt.Errorf("cannot get account next index:%w", err)
	}

	options := types.SignatureOptions{
		BlockHash:          GenesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        GenesisHash,
		Nonce:              nonce,
		SpecVersion:        Rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(defaultTip),
		AppID:              types.NewUCompactFromUInt(uint64(AppId)),
		TransactionVersion: Rv.TransactionVersion,
	}

	err = ext.Sign(KeyringPair, options)
	if err != nil {
		return nil, fmt.Errorf("cannot sign:%w", err)
	}

	// Send the extrinsic
	sub, err := Api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return nil, fmt.Errorf("cannot submit extrinsic:%w", err)
	}

	defer sub.Unsubscribe()
	timeout := time.After(time.Duration(Config.Timeout) * time.Second)
	var blockHash types.Hash
out:
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				log.Infof("📥 Submit data extrinsic included in block %v", status.AsInBlock.Hex())
			}
			if status.IsFinalized {
				blockHash = status.AsFinalized
				break out
			} else if status.IsDropped {
				return nil, fmt.Errorf("❌ Extrinsic dropped")
			} else if status.IsUsurped {
				return nil, fmt.Errorf("❌ Extrinsic usurped")
			} else if status.IsRetracted {
				return nil, fmt.Errorf("❌ Extrinsic retracted")
			} else if status.IsInvalid {
				return nil, fmt.Errorf("❌ Extrinsic invalid")
			}
		case <-timeout:
			return nil, fmt.Errorf("⌛️ Timeout of %d seconds reached without getting finalized status for extrinsic", Config.Timeout)
		}
	}

	log.Infof("✅ Data submitted by sequencer:%d bytes against AppID %v sent with hash %#x", len(txData), AppId, blockHash)

	var dataProof DataProof
	var batchHash [32]byte

	h := sha3.NewLegacyKeccak256()
	h.Write(txData)
	h.Sum(batchHash[:0])

	block, err := Api.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get block:%w", err)
	}

	for i := 1; i <= len(block.Block.Extrinsics); i++ {
		resp, err := http.Post("https://goldberg.avail.tools/api", "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"kate_queryDataProof\",\"params\":[%d, \"%#x\"]}", i, blockHash)))
		if err != nil {
			return nil, fmt.Errorf("cannot post query request:%v", err)
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("cannot read body:%v", err)
		}

		var dataProofResp DataProofRPCResponse
		if err := json.Unmarshal(data, &dataProofResp); err != nil {
			return nil, fmt.Errorf("cannot unmarshal data proof:%v", err)
		}

		if dataProofResp.Result.Leaf == fmt.Sprintf("%#x", batchHash) {
			dataProof = dataProofResp.Result
			break
		}
	}

	log.Infof("💿 received data proof:%+v", dataProof)
	var batchDAData availTypes.BatchDAData
	batchDAData.Proof = dataProof.Proof
	batchDAData.Width = dataProof.NumberOfLeaves
	batchDAData.LeafIndex = dataProof.LeafIndex

	header, err := Api.RPC.Chain.GetHeader(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get header:%+v", err)
	}

	batchDAData.BlockNumber = uint(header.Number)

	//nolint:errcheck
	GetData(uint64(header.Number), dataProof.LeafIndex)
	log.Infof("🟢 prepared DA data:%+v", batchDAData)
	return &batchDAData, nil
}

// DispatchDataRoot dispatches the data root to the avail DA
func DispatchDataRoot(blockNumber uint64) error {
	blockHash, err := Api.RPC.Chain.GetBlockHash(blockNumber)
	if err != nil {
		return fmt.Errorf("cannot get runtime version:%w", err)
	}

	header, err := Api.RPC.Chain.GetHeader(blockHash)
	if err != nil {
		return fmt.Errorf("cannot get header:%+v", err)
	}

	dispatchDataRootCall, err := types.NewCall(Meta, "NomadDABridge.try_dispatch_data_root", DestinationDomain, DestinationAddress, header)

	if err != nil {
		return fmt.Errorf("cannot create new call:%+v", err)
	}

	dispatchDataRootExt := types.NewExtrinsic(dispatchDataRootCall)

	nonce, err := GetAccountNextIndex()
	if err != nil {
		return fmt.Errorf("cannot get account next index:%w", err)
	}

	options := types.SignatureOptions{
		BlockHash:          GenesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        GenesisHash,
		Nonce:              nonce,
		SpecVersion:        Rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(defaultTip),
		AppID:              types.NewUCompactFromUInt(0),
		TransactionVersion: Rv.TransactionVersion,
	}

	err = dispatchDataRootExt.Sign(KeyringPair, options)
	if err != nil {
		return fmt.Errorf("cannot sign:%w", err)
	}

	sub, err := Api.RPC.Author.SubmitAndWatchExtrinsic(dispatchDataRootExt)
	if err != nil {
		return fmt.Errorf("cannot dispatch data root extrinsic:%+v", err)
	}
	defer sub.Unsubscribe()
	timeout := time.After(time.Duration(Config.Timeout) * time.Second)
out:
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				log.Infof("📥 Dispatch data root extrinsic included in block %v", status.AsInBlock.Hex())
				break out
			} else if status.IsDropped {
				return fmt.Errorf("❌ Extrinsic dropped")
			} else if status.IsUsurped {
				return fmt.Errorf("❌ Extrinsic usurped")
			} else if status.IsInvalid {
				return fmt.Errorf("❌ Extrinsic invalid")
			} else if status.IsRetracted {
				return fmt.Errorf("❌ Extrinsic retracted")
			}
		case <-timeout:
			return fmt.Errorf("⌛️ Timeout of %d seconds reached without getting finalized status for dispatch data root extrinsic", Config.Timeout)
		}
	}

	log.Infof("✅ Data root dispatched by sequencer sent in block %#x\n", blockHash)

	return nil
}

// GetData gets the data from the avail DA block
func GetData(blockNumber uint64, index uint) ([]byte, error) {
	blockHash, err := Api.RPC.Chain.GetBlockHash(uint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("cannot get block hash:%w", err)
	}

	block, err := Api.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return nil, fmt.Errorf("cannot get block:%w", err)
	}

	var data [][]byte
	for _, ext := range block.Block.Extrinsics {
		if ext.Method.CallIndex.SectionIndex == 29 && ext.Method.CallIndex.MethodIndex == 1 {
			data = append(data, ext.Method.Args[2:])
		}
	}

	return data[index], nil
}

// GetAccountNextIndex gets the next index of the account from avail DA
func GetAccountNextIndex() (types.UCompact, error) {
	resp, err := http.Post("https://goldberg.avail.tools/api", "application/json", strings.NewReader(fmt.Sprintf("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"system_accountNextIndex\",\"params\":[\"%v\"]}", KeyringPair.Address)))
	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot post query request:%v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot read body:%v", err)
	}

	var accountNextIndex AccountNextIndexRPCResponse

	if err := json.Unmarshal(data, &accountNextIndex); err != nil {
		return types.NewUCompactFromUInt(0), fmt.Errorf("cannot unmarshal account next index:%v", err)
	}

	return types.NewUCompactFromUInt(uint64(accountNextIndex.Result)), nil
}
