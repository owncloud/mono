package svc

import (
	"github.com/owncloud/ocis-pkg/v2/log"
)

// NewLogging returns a service that logs messages.
func NewLogging(next Service, logger log.Logger) Service {
	return Service{
		manager: next.manager,
		config:  next.config,
	}
}
