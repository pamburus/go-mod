// Package portalloc provides a way to allocate a port number that is available for use.
package portalloc

import (
	"fmt"
	"net"
)

// New returns a new port number that is available for use.
func New() (uint16, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()

	actualAddr := l.Addr()
	tcpAddr, ok := actualAddr.(*net.TCPAddr)
	if !ok {
		return 0, fmt.Errorf("listener address has unexpected type %T (expected %T)", actualAddr, addr)
	}

	return tcpAddr.AddrPort().Port(), nil
}
