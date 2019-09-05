package main

import (
	"os"

	"github.com/owncloud/reva-ocs/pkg/command"
)

func main() {
	if err := command.Root().Execute(); err != nil {
		os.Exit(1)
	}
}