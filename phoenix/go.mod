module github.com/owncloud/mono/phoenix

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/ocagent v0.6.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/UnnoTed/fileb0x v1.1.4
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/oklog/run v1.0.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/owncloud/mono/ocis-pkg v0.0.0-00010101000000-000000000000
	github.com/owncloud/ocis-phoenix v0.13.0 // indirect
	github.com/owncloud/ocis-pkg v1.3.0 // indirect
	github.com/restic/calens v0.2.0
	github.com/spf13/viper v1.6.3
	go.opencensus.io v0.22.4
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
)

replace (
	github.com/owncloud/mono/ocis-pkg => ../ocis-pkg
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
