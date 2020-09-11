package grpc

import (
	"github.com/owncloud/mono/ocis-pkg/service/grpc"
	"github.com/owncloud/mono/settings/pkg/proto/v0"
	svc "github.com/owncloud/mono/settings/pkg/service/v0"
	"github.com/owncloud/mono/settings/pkg/version"
)

// Server initializes a new go-micro service ready to run
func Server(opts ...Option) grpc.Service {
	options := newOptions(opts...)

	service := grpc.NewService(
		grpc.Logger(options.Logger),
		grpc.Name(options.Name),
		grpc.Version(version.String),
		grpc.Address(options.Config.GRPC.Addr),
		grpc.Namespace(options.Config.GRPC.Namespace),
		grpc.Context(options.Context),
		grpc.Flags(options.Flags...),
	)

	handle := svc.NewService(options.Config, options.Logger)
	if err := proto.RegisterBundleServiceHandler(service.Server(), handle); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register Bundle service handler")
	}
	if err := proto.RegisterValueServiceHandler(service.Server(), handle); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register Value service handler")
	}
	if err := proto.RegisterRoleServiceHandler(service.Server(), handle); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register Role service handler")
	}
	if err := proto.RegisterPermissionServiceHandler(service.Server(), handle); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register Permission service handler")
	}

	service.Init()
	return service
}
