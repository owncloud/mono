module github.com/owncloud/ocis/konnectd

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/ocagent v0.6.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/UnnoTed/fileb0x v1.1.4
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/gorilla/mux v1.7.4
	github.com/haya14busa/goverage v0.0.0-20180129164344-eec3514a20b5
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/oklog/run v1.0.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/owncloud/flaex v0.2.0
	github.com/owncloud/ocis/ocis-pkg v0.0.0-20200911111357-8d02528b0f85
	github.com/restic/calens v0.2.0
	github.com/rs/zerolog v1.19.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/viper v1.6.3
	go.opencensus.io v0.22.4
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	stash.kopano.io/kc/konnect v0.30.0
	stash.kopano.io/kgol/rndm v1.1.0
)

replace (
	github.com/owncloud/ocis/konnectd => ../konnectd
	github.com/owncloud/ocis/ocis-pkg => ../ocis-pkg
	github.com/owncloud/ocis/settings => ../settings
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
