// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package digest

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/hash"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/pkg/errors"
)

const (
	TX_ID_SIZE_BYTES = 8 + 32
)

func CalcTxHash(transaction *protocol.Transaction) primitives.Sha256 {
	return hash.CalcSha256(transaction.Raw())
}

func CalcTxId(transaction *protocol.Transaction) []byte {
	return GenerateTxId(CalcTxHash(transaction), transaction.Timestamp())
}

func GenerateTxId(txHash primitives.Sha256, txTimestamp primitives.TimestampNano) []byte {
	res := make([]byte, TX_ID_SIZE_BYTES)
	membuffers.WriteUint64(res, uint64(txTimestamp))
	copy(res[8:], txHash)

	return res
}

func ExtractTxId(txId []byte) (txHash primitives.Sha256, txTimestamp primitives.TimestampNano, err error) {
	if len(txId) != TX_ID_SIZE_BYTES {
		err = errors.Errorf("txid has invalid length %d", len(txId))
		return
	}
	txTimestamp = primitives.TimestampNano(membuffers.GetUint64(txId))
	txHash = txId[8:]
	return
}
