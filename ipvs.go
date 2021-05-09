package main

import (
	"sync"

	"github.com/kwanhur/ipvs"
)

var ipvsMutex sync.Mutex

type IPVS struct {
	handler *ipvs.Handle
}

// NewIPVS return ipvs wrapper
func NewIPVS() (*IPVS, error) {
	ipvsMutex.Lock()
	defer ipvsMutex.Unlock()

	if handler, err := ipvs.New(); err != nil {
		return nil, err
	} else {
		return &IPVS{
			handler: handler,
		}, nil
	}
}

// Close close ipvs netlink socket handler
func (s *IPVS) Close() {
	if s.handler != nil {
		s.handler.Close()
	}
}

// Info return ipvs infomation, include version number and connection table size
func (s *IPVS) Info() (*ipvs.Info, error) {
	return s.handler.GetInfo()
}
