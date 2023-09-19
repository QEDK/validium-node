package types

import (
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

// Sequence represents an operation sent to the PoE smart contract to be
// processed.
type Sequence struct {
	GlobalExitRoot, StateRoot, LocalExitRoot common.Hash //
	AccInputHash                             common.Hash // 1024
	Timestamp                                int64       //64
	BatchHash                                [32]byte
	IsSequenceTooBig                         bool   // 8
	BatchNumber                              uint64 // 64
	ForcedBatchTimestamp                     int64  // 64
	DABlockNumber                              uint32
	DAProof                                    [][32]byte
	DAWidth                                    big.Int
	DAIndex                                    big.Int
}

// IsEmpty checks is sequence struct is empty
func (s Sequence) IsEmpty() bool {
	return reflect.DeepEqual(s, Sequence{})
}
