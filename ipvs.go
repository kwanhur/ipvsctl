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
	Handler *ipvs.Handle
}

// NewIPVS return ipvs wrapper
func NewIPVS() (*IPVS, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if handler, err := ipvs.New(""); err != nil {
		return nil, err
	} else {
		return &IPVS{
			Handler: handler,
		}, nil
	}
}

// Close close ipvs netlink socket Handler
func (s *IPVS) Close() {
	if s.Handler != nil {
		s.Handler.Close()
	}
}
