// Copyright 2022 kwanhur
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
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

func (s IPVS) ExistService(svc *ipvs.Service) bool {
	return s.handler.IsServicePresent(svc)
}

// DelService delete ipvs service
func (s IPVS) DelService(svc *ipvs.Service) error {
	return s.handler.DelService(svc)
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
