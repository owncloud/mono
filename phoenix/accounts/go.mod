module github.com/owncloud/mono/accounts

go 1.13

require (
	github.com/CiscoM31/godata v0.0.0-20191007193734-c2c4ebb1b415
	github.com/blevesearch/bleve v1.0.9
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/mennanov/fieldmask-utils v0.3.2
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/oklog/run v1.1.0
	github.com/olekukonko/tablewriter v0.0.4
	github.com/owncloud/mono/ocis-pkg v0.0.0-20200911111357-8d02528b0f85
	github.com/owncloud/mono/settings v0.0.0-20200911111357-8d02528b0f85
	github.com/owncloud/ocis v1.0.0-rc1 // indirect
	github.com/restic/calens v0.2.0
	github.com/rs/zerolog v1.19.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.6.1
	github.com/tredoe/osutil v1.0.5
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/genproto v0.0.0-20200624020401-64a14ca9d1ad
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/owncloud/mono/ocis-pkg => ../ocis-pkg
	github.com/owncloud/mono/settings => ../settings
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
