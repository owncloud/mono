module github.com/owncloud/mono/ocis

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/ocagent v0.7.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/UnnoTed/fileb0x v1.1.4
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/micro/v2 v2.8.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/owncloud/flaex v0.2.0
	github.com/owncloud/mono/accounts v0.4.2-0.20200901074457-6a27781a2741
	github.com/owncloud/mono/ocis-pkg v0.0.0-20200911132118-5feb827c3e6c
	github.com/owncloud/mono/proxy v0.0.0-00010101000000-000000000000
	github.com/owncloud/mono/settings v0.0.0-20200911111357-8d02528b0f85
	github.com/owncloud/mono/thumbnails v0.1.6
	github.com/owncloud/mono/webdav v0.0.0-00010101000000-000000000000
	github.com/owncloud/ocis v1.0.0-rc1
	github.com/owncloud/ocis-glauth v0.5.1-0.20200909064150-0735ec933777
	github.com/owncloud/ocis-graph v0.0.0-20200318175820-9a5a6e029db7
	github.com/owncloud/ocis-graph-explorer v0.0.0-20200210111049-017eeb40dc0c
	github.com/owncloud/ocis-hello v0.1.0-alpha1.0.20200828085053-37fcf3c8f853
	github.com/owncloud/ocis-konnectd v0.3.2
	github.com/owncloud/ocis-migration v0.2.0 // indirect
	github.com/owncloud/ocis-ocs v0.3.1
	github.com/owncloud/ocis-phoenix v0.13.1-0.20200910090722-8e3ed0e966a5
	github.com/owncloud/ocis-proxy v0.7.1-0.20200907105449-201b9a652685 // indirect
	github.com/owncloud/ocis-reva v0.13.1-0.20200910085648-26465bbdcf46
	github.com/owncloud/ocis-store v0.1.1
	github.com/refs/pman v0.0.0-20200701173654-f05b8833071a
	github.com/restic/calens v0.2.0
	go.opencensus.io v0.22.4
)

replace (
	github.com/gomodule/redigo => github.com/gomodule/redigo v1.8.2
	github.com/owncloud/mono/accounts => ../accounts
	github.com/owncloud/mono/ocis-pkg => ../ocis-pkg
	github.com/owncloud/mono/proxy => ../proxy
	github.com/owncloud/mono/settings => ../settings
	github.com/owncloud/mono/thumbnails => ../thumbnails
	github.com/owncloud/mono/webdav => ../webdav
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
