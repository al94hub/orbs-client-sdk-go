// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package orbs

import (
	"github.com/orbs-network/orbs-client-sdk-go/crypto/digest"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/encoding"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/keys"
)

type OrbsAccount struct {
	PublicKey  []byte
	PrivateKey []byte
	Address    string // hex string starting with 0x
}

func CreateAccount() (*OrbsAccount, error) {
	keyPair, err := keys.GenerateEd25519Key()
	if err != nil {
		return nil, err
	}

	rawAddress, err := digest.CalcClientAddressOfEd25519PublicKey(keyPair.PublicKey())
	if err != nil {
		return nil, err
	}

	return &OrbsAccount{
		PublicKey:  keyPair.PublicKey(),
		PrivateKey: keyPair.PrivateKey(),
		Address:    BytesToAddress(rawAddress),
	}, nil
}

func (oa *OrbsAccount) AddressAsBytes() []byte {
	return AddressToBytes(oa.Address)
}

func AddressToBytes(address string) []byte {
	rawAddress, err := encoding.DecodeHex(address)
	if err != nil {
		return nil
	}
	return rawAddress
}

func AddressValidate(address string) error {
	_, err := encoding.DecodeHex(address)
	return err
}

func BytesToAddress(rawAddress []byte) string {
	return encoding.EncodeHex(rawAddress)
}
