Enhancement: Add option to disable TLS

Can be used to disable TLS when the ocis-proxy is behind an
TLS-Terminating reverse proxy.

env PROXY_TLS=false or --tls=false

https://github.com/owncloud/mono/proxy/issues/71
https://github.com/owncloud/mono/proxy/pull/72
