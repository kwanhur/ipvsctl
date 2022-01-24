package main

import (
	"sync"

	"github.com/kwanhur/ipvs"
)

var mutex sync.Mutex

type IPVS struct {
	handler *ipvs.Handle
}

// NewIPVS return ipvs wrapper
func NewIPVS() (*IPVS, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if handler, err := ipvs.New(""); err != nil {
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

// Info return ipvs information, include version number and connection table size
func (s *IPVS) Info() (*ipvs.Info, error) {
	return s.handler.GetInfo()
}

// AddService add ipvs service
func (s *IPVS) AddService(svc *ipvs.Service) error {
	return s.handler.NewService(svc)
}

// Flush clear out ipvs rules
func (s *IPVS) Flush() error {
	return s.handler.Flush()
}

// Config return ipvs configuration
func (s *IPVS) Config() (*ipvs.Config, error) {
	return s.handler.GetConfig()
}

// SetConfig set ipvs configuration
func (s *IPVS) SetConfig(c *ipvs.Config) error {
	return s.handler.SetConfig(c)
}
