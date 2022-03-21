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
	"regexp"
	"strconv"
	"strings"
)

// Connection ip_vs connection
type Connection struct {
	Protocol    string
	ClientIP    string
	ClientPort  uint16
	VirtualIP   string
	VirtualPort uint16
	LocalIP     string
	LocalPort   uint16
	ServerIP    string
	ServerPort  uint16
	State       string
	Expire      uint64
	PEName      string
}

// ExpireString convert to format(minute:second)
func (c *Connection) ExpireString() string {
	if c.Expire <= 0 {
		return fmt.Sprintf("%d", c.Expire)
	}

	secs := c.Expire % 60
	mins := c.Expire / 60

	return fmt.Sprintf("%d:%d", mins, secs)
}

func toConnections(body []byte) []*Connection {
	var reg = regexp.MustCompile(`\s+`)
	var conns []*Connection

	for _, line := range strings.Split(string(body), "\n") {
		line = reg.ReplaceAllString(line, "\t")
		lines := strings.Split(line, "\t")
		if len(lines) < 9 || lines[0] != "Pro" {
			continue
		}

		protocol := lines[0]
		clientIP := ipAddress(lines[1])
		clientPort, _ := strconv.ParseUint(lines[2], 16, 0)
		virtualIP := ipAddress(lines[3])
		virtualPort, _ := strconv.ParseUint(lines[4], 16, 0)

		var (
			localIP    = ""
			localPort  = uint64(0)
			serverIP   = ""
			serverPort = uint64(0)
			state      = ""
			expire     = uint64(0)
			pename     = ""
		)

		if len(lines) >= 11 {
			localIP = ipAddress(lines[5])
			localPort, _ = strconv.ParseUint(lines[6], 16, 0)
			serverIP = ipAddress(lines[7])
			serverPort, _ = strconv.ParseUint(lines[8], 16, 0)
			state = lines[9]
			expire, _ = strconv.ParseUint(lines[10], 10, 0)
			if len(lines) >= 12 {
				pename = lines[11]
			}
		} else if len(lines) >= 9 {
			serverIP = ipAddress(lines[5])
			serverPort, _ = strconv.ParseUint(lines[6], 16, 0)
			state = lines[7]
			expire, _ = strconv.ParseUint(lines[8], 10, 0)
			if len(lines) >= 10 {
				pename = lines[9]
			}
		}

		conn := &Connection{
			Protocol:    protocol,
			ClientIP:    clientIP,
			ClientPort:  uint16(clientPort),
			VirtualIP:   virtualIP,
			VirtualPort: uint16(virtualPort),
			LocalIP:     localIP,
			LocalPort:   uint16(localPort),
			ServerIP:    serverIP,
			ServerPort:  uint16(serverPort),
			State:       state,
			Expire:      expire,
			PEName:      pename,
		}

		conns = append(conns, conn)
	}

	return conns
}
