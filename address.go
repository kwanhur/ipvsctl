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
	"fmt"
	"net"
	"strconv"
)

func ipAddress(addr string) string {
	if len(addr) == 8 {
		dot1, _ := strconv.ParseUint(addr[0:2], 16, 0)
		dot2, _ := strconv.ParseUint(addr[2:4], 16, 0)
		dot3, _ := strconv.ParseUint(addr[4:6], 16, 0)
		dot4, _ := strconv.ParseUint(addr[6:8], 16, 0)

		return fmt.Sprintf("%d.%d.%d.%d", dot1, dot2, dot3, dot4)
	}

	return net.ParseIP(addr).String()
}
