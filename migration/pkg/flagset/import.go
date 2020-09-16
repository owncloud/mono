package flagset

import (
	"github.com/micro/cli/v2"
	"github.com/owncloud/ocis-migration/pkg/config"
)

// ImportWithConfig applies cfg to the import-command
func ImportWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "import-path",
			Usage:    "Path to exported user-directory",
			Value:    "",
			Required: true,
			EnvVars:  []string{"MIGRATION_IMPORT_PATH"},
		},
		&cli.StringFlag{
			Name:    "jwt-secret",
			Value:   "Pive-Fumkiu4",
			Usage:   "Used to create JWT to talk to reva, should equal reva's jwt-secret",
			EnvVars: []string{"MIGRATION_JWT_SECRET"},
		},
	}
}
