package avail

import (
	"fmt"

	"github.com/0xPolygonHermez/zkevm-node/avail/internal/config"
	evmTypes "github.com/ethereum/go-ethereum/core/types"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// The following example shows how submit data blob and track transaction status
func PostData(txData []evmTypes.Transaction) {
	var config config.Config

	err := config.GetConfig("/app/avail-config.json")
	if err != nil {
		panic(fmt.Sprintf("cannot get config:%v", err))
	}

	api, err := gsrpc.NewSubstrateAPI(config.ApiURL)
	if err != nil {
		panic(fmt.Sprintf("cannot create api:%v", err))
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(fmt.Sprintf("cannot get metadata:%v", err))
	}

	var subData []byte
	for i := 0; i < len(txData); i++ {
		bytes, err := txData[i].MarshalBinary()
		if err != nil {
			panic(fmt.Sprintf("invalid tx from finalizer:%v", err))
		}
		subData = append(subData, bytes...)
	}
	fmt.Println("⚡️ Submitting data to Avail...")
	appID := 0

	// if app id is greater than 0 then it must be created before submitting data
	if config.AppID != 0 {
		appID = config.AppID
	}

	newCall, err := types.NewCall(meta, "DataAvailability.submit_data", types.NewBytes(subData))
	if err != nil {
		panic(fmt.Sprintf("cannot create new call:%v", err))
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(newCall)

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		panic(fmt.Sprintf("cannot get block hash:%v", err))
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		panic(fmt.Sprintf("cannot get latest runtime version:%v", err))
	}

	keyringPair, err := signature.KeyringPairFromSecret(config.Seed, 42)
	if err != nil {
		panic(fmt.Sprintf("cannot create KeyPair:%v", err))
	}

	key, err := types.CreateStorageKey(meta, "System", "Account", keyringPair.PublicKey)
	if err != nil {
		panic(fmt.Sprintf("cannot create storage key:%w", err))
	}

	var accountInfo types.AccountInfo
	ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		panic(fmt.Sprintf("cannot get latest storage:%v", err))
	}

	nonce := uint32(accountInfo.Nonce)
	options := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(100),
		AppID:              types.NewUCompactFromUInt(uint64(appID)),
		TransactionVersion: rv.TransactionVersion,
	}

	err = ext.Sign(keyringPair, options)
	if err != nil {
		panic(fmt.Sprintf("cannot sign:%v", err))
	}

	// Send the extrinsic
	hash, err := api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		panic(fmt.Sprintf("cannot submit extrinsic:%v", err))
	}

	fmt.Printf("Data submitted by sequencer: %v against appID %v  sent with hash %#x\n", subData, appID, hash)
}
