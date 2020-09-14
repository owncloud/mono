Bugfix: Fix file descriptor leak

Only use a single instance of go-micro's GRPC client as it already
does connection pooling. This prevents connection and file descriptor leaks.

https://github.com/owncloud/mono/accounts/issues/79
https://github.com/owncloud/mono/ocs/pull/29
