package main

import (
	"os"

	"github.com/vlototskyi/ergo-blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
