package encoding

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-client-sdk-go/crypto/hash"
	"github.com/pkg/errors"
	"strings"
)

func EncodeHex(data []byte) string { // EIP-55 'complaint', but using sha2 and not sha3
	result := []byte(hex.EncodeToString(data))
	hashed := hash.CalcSha256(data)

	for i := 0; i < len(result); i++ {
		hashByte := hashed[(i/2)%hash.SHA256_HASH_SIZE_BYTES]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}

		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}

	return "0x" + string(result)
}

// on decode error (eg. non hex character in str) returns zero_value, error
// on checksum failure returns decoded_value, error (so users could warn about checksum but still use the decoded)
// if all is lower or upper then the checksum check is ignored (as the checksum was probably not taken into account)
func DecodeHex(str string) ([]byte, error) {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}

	data, err := hex.DecodeString(str)
	if err != nil {
		return nil, errors.Wrap(err, "invalid hex string")
	}

	encoded := EncodeHex(data)
	if encoded[2:] != str {
		// checksum error, we will allow if the source is in uniform case (all lower/upper)
		if strings.ToUpper(str) == str || strings.ToLower(str) == str {
			return data, nil
		} else {
			return data, errors.New("invalid checksum")
		}
	}

	return data, nil
}
