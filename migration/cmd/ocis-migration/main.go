package main

import (
	"os"

	"github.com/owncloud/mono/migration/pkg/command"
)

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
