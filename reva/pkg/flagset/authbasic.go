package flagset

import (
	"github.com/micro/cli/v2"
	"github.com/owncloud/ocis-reva/pkg/config"
)

// AuthBasicWithConfig applies cfg to the root flagset
func AuthBasicWithConfig(cfg *config.Config) []cli.Flag {
	flags := []cli.Flag{

		// debug ports are the odd ports
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:9147",
			Usage:       "Address to bind debug server",
			EnvVars:     []string{"REVA_AUTH_BASIC_DEBUG_ADDR"},
			Destination: &cfg.Reva.AuthBasic.DebugAddr,
		},

		// Auth

		&cli.StringFlag{
			Name:        "auth-driver",
			Value:       "ldap",
			Usage:       "auth driver: 'demo', 'json' or 'ldap'",
			EnvVars:     []string{"REVA_AUTH_DRIVER"},
			Destination: &cfg.Reva.AuthProvider.Driver,
		},
		&cli.StringFlag{
			Name:        "auth-json",
			Value:       "",
			Usage:       "Path to users.json file",
			EnvVars:     []string{"REVA_AUTH_JSON"},
			Destination: &cfg.Reva.AuthProvider.JSON,
		},

		// Services

		// AuthBasic

		&cli.StringFlag{
			Name:        "network",
			Value:       "tcp",
			Usage:       "Network to use for the reva auth-basic service, can be 'tcp', 'udp' or 'unix'",
			EnvVars:     []string{"REVA_AUTH_BASIC_NETWORK"},
			Destination: &cfg.Reva.AuthBasic.Network,
		},
		&cli.StringFlag{
			Name:        "protocol",
			Value:       "grpc",
			Usage:       "protocol for reva service, can be 'http' or 'grpc'",
			EnvVars:     []string{"REVA_AUTH_BASIC_PROTOCOL"},
			Destination: &cfg.Reva.AuthBasic.Protocol,
		},
		&cli.StringFlag{
			Name:        "addr",
			Value:       "0.0.0.0:9146",
			Usage:       "Address to bind reva service",
			EnvVars:     []string{"REVA_AUTH_BASIC_ADDR"},
			Destination: &cfg.Reva.AuthBasic.Addr,
		},
		&cli.StringFlag{
			Name:        "url",
			Value:       "localhost:9146",
			Usage:       "URL to use for the reva service",
			EnvVars:     []string{"REVA_AUTH_BASIC_URL"},
			Destination: &cfg.Reva.AuthBasic.URL,
		},
		&cli.StringSliceFlag{
			Name:    "service",
			Value:   cli.NewStringSlice("authprovider"),
			Usage:   "--service authprovider [--service otherservice]",
			EnvVars: []string{"REVA_AUTH_BASIC_SERVICES"},
		},
	}

	flags = append(flags, TracingWithConfig(cfg)...)
	flags = append(flags, DebugWithConfig(cfg)...)
	flags = append(flags, SecretWithConfig(cfg)...)
	flags = append(flags, LDAPWithConfig(cfg)...)

	return flags
}
