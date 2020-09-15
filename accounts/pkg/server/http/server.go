package http

import (
	"github.com/go-chi/chi"
	mclient "github.com/micro/go-micro/v2/client"
	"github.com/owncloud/mono/accounts/pkg/assets"
	"github.com/owncloud/mono/accounts/pkg/proto/v0"
	svc "github.com/owncloud/mono/accounts/pkg/service/v0"
	"github.com/owncloud/mono/accounts/pkg/version"
	"github.com/owncloud/mono/ocis-pkg/account"
	"github.com/owncloud/mono/ocis-pkg/middleware"
	"github.com/owncloud/mono/ocis-pkg/roles"
	"github.com/owncloud/mono/ocis-pkg/service/http"
	settings "github.com/owncloud/mono/settings/pkg/proto/v0"
)

// Server initializes the http service and server.
func Server(opts ...Option) http.Service {
	options := newOptions(opts...)
	handler := options.Handler

	service := http.NewService(
		http.Logger(options.Logger),
		http.Name(options.Name),
		http.Version(version.String),
		http.Address(options.Config.HTTP.Addr),
		http.Namespace(options.Config.HTTP.Namespace),
		http.Context(options.Context),
		http.Flags(options.Flags...),
	)

	mux := chi.NewMux()

	mux.Use(middleware.RealIP)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Cache)
	mux.Use(middleware.Cors)
	mux.Use(middleware.Secure)
	mux.Use(middleware.ExtractAccountUUID(
		account.Logger(options.Logger),
		account.JWTSecret(options.Config.TokenManager.JWTSecret)),
	)

	mux.Use(middleware.Version(
		options.Name,
		version.String,
	))

	mux.Use(middleware.Logger(
		options.Logger,
	))

	mux.Use(middleware.Static(
		options.Config.HTTP.Root,
		assets.New(
			assets.Logger(options.Logger),
			assets.Config(options.Config),
		),
	))

	mux.Route(options.Config.HTTP.Root, func(r chi.Router) {
		proto.RegisterAccountsServiceWeb(r, handler)
		proto.RegisterGroupsServiceWeb(r, handler)
	})

	service.Handle(
		"/",
		mux,
	)

	if err := service.Init(); err != nil {
		panic(err)
	}
	return service
}
