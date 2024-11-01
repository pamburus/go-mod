package sqltest

import (
	"context"
	"encoding/base32"
	"encoding/binary"
	"hash"
	"net"
	"os"
	"strings"
	"time"
)

func freePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}

func hashBase32(hash hash.Hash, data ...[]byte) string {
	for _, data := range data {
		binary.Write(hash, binary.LittleEndian, int64(len(data)))
		hash.Write(data)
	}

	return strings.ToLower(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hash.Sum(nil)))
}

func envFlag(name string) bool {
	if val, ok := os.LookupEnv(name); ok {
		return enabled[strings.ToLower(val)]
	}

	return false
}

func sleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// ---

var enabled = map[string]bool{
	"on":   true,
	"yes":  true,
	"1":    true,
	"true": true,
}
