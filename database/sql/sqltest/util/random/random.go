package random

import (
	"encoding/binary"
	"hash/fnv"
	"math/rand/v2"

	"github.com/pamburus/go-mod/database/sql/sqltest/util/hashstr"
)

func Password(source rand.Source) string {
	return hashstr.New(fnv.New64(), binary.AppendUvarint(nil, source.Uint64()))
}
