Enhancement: Remove paths from Makefile

We have a variable for the proto files path in our Makefile, but were not using
it. Changed the Makefile to use the PROTO_SRC variable where possible.

https://github.com/owncloud/mono/settings/pull/33
