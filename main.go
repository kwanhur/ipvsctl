//go:build linux
// +build linux

package main

import (
	"github.com/urfave/cli/v2"
)

var (
	Version  string
	CommitID string
	Built    string
)

func main() {
	app := cli.NewApp()
	app.Usage = "IP Virtual Server controller"
	app.Version = Version
	app.Description = "ipvs controller communicate with ip_vs kernel module"
	app.Authors = Authors()

	opr := NewOperator()
	cli.VersionPrinter = opr.ShowVersion()

	app.Commands = opr.Commands()

	app.RunAndExitOnError()
}
