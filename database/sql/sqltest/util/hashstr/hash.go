// Package hashstr provides a simple way to generate a base32 encoded hash of some data.
package hashstr

import (
	"encoding/base32"
	"encoding/binary"
	"hash"
	"strings"
)

// New returns a base32 encoded hash of the given data.
func New(hash hash.Hash, data ...[]byte) string {
	for _, data := range data {
		binary.Write(hash, binary.LittleEndian, int64(len(data)))
		hash.Write(data)
	}

	return strings.ToLower(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hash.Sum(nil)))
}
