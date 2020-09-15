package service

import (
	"github.com/owncloud/mono/accounts/pkg/config"
	"github.com/owncloud/mono/ocis-pkg/log"
	"github.com/owncloud/mono/ocis-pkg/roles"
	settings "github.com/owncloud/mono/settings/pkg/proto/v0"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Logger      log.Logger
	Config      *config.Config
	RoleService settings.RoleService
	RoleManager *roles.Manager
}

func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Logger provides a function to set the Logger option.
func Logger(val log.Logger) Option {
	return func(o *Options) {
		o.Logger = val
	}
}

// Config provides a function to set the Config option.
func Config(val *config.Config) Option {
	return func(o *Options) {
		o.Config = val
	}
}

// RoleService provides a function to set the RoleService option.
func RoleService(val settings.RoleService) Option {
	return func(o *Options) {
		o.RoleService = val
	}
}

// RoleManager provides a function to set the RoleManager option.
func RoleManager(val *roles.Manager) Option {
	return func(o *Options) {
		o.RoleManager = val
	}
}
