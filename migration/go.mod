module github.com/owncloud/ocis/migration

go 1.13

require (
	github.com/UnnoTed/fileb0x v1.1.4
	github.com/cs3org/go-cs3apis v0.0.0-20200810113633-b00aca449666
	github.com/cs3org/reva v1.2.1-0.20200826162318-c0f54e1f37ea
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/owncloud/ocis/accounts v0.0.0-00010101000000-000000000000
	github.com/owncloud/ocis-pkg/v2 v2.4.1-0.20200902134813-1e87c6173ada
	github.com/restic/calens v0.2.0
	github.com/spf13/viper v1.7.0
	google.golang.org/grpc v1.31.1
)

replace (
	github.com/owncloud/ocis/accounts => ../accounts
	github.com/owncloud/ocis/ocis-pkg => ../ocis-pkg
)
