package pex

import (
	"fmt"

	"github.com/tendermint/tendermint/p2p"
)

type ErrAddrBookNonRoutable struct {
	Addr *p2p.NetAddress
}

func (err ErrAddrBookNonRoutable) Error() string {
	return fmt.Sprintf("Cannot add non-routable address %v", err.Addr)
}

type ErrAddrBookSelf struct {
	Addr *p2p.NetAddress
}

func (err ErrAddrBookSelf) Error() string {
	return fmt.Sprintf("Cannot add ourselves with address %v", err.Addr)
}

type ErrAddrBookNilAddr struct {
	Addr *p2p.NetAddress
	Src  *p2p.NetAddress
}

func (err ErrAddrBookNilAddr) Error() string {
	return fmt.Sprintf("Cannot add a nil address. Got (addr, src) = (%v, %v)", err.Addr, err.Src)
}

type ErrAddrBookFull struct {
	Addr *p2p.NetAddress
	Size int
}

func (err ErrAddrBookFull) Error() string {
	return fmt.Sprintf("Can't add new address (%v), addr book is full (%d)", err.Addr, err.Size)
}