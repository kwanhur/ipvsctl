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
//go:build linux
// +build linux

package main

import (
	"github.com/urfave/cli/v2"
)

var (
	// Version ipvsctl version
	Version string
	// CommitID ipvsctl git commit id
	CommitID string
	// Built build ipvsctl date
	Built string
)

func main() {
	app := cli.NewApp()
	app.Usage = "IP Virtual Server controller"
	app.Version = Version
	app.Description = "A modern Linux Virtual Server controller"
	app.Authors = Authors()

	opr := NewOperator()
	cli.VersionPrinter = opr.ShowVersion()

	app.Commands = opr.Commands()

	app.RunAndExitOnError()
}
