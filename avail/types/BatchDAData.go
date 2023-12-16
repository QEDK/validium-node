// SPDX-License-Identifier: Apache-2.0
package types

import (
	"reflect"
)

// BatchDAData represents the data of a batch
type BatchDAData struct {
	BlockNumber uint
	Proof       []string `json:"proof"`
	Width       uint     `json:"number_of_leaves"`
	LeafIndex   uint     `json:"leaf_index"`
}

// IsEmpty returns true if the BatchDAData is empty
func (b BatchDAData) IsEmpty() bool {
	return reflect.DeepEqual(b, BatchDAData{})
}
