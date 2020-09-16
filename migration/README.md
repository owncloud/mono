# ownCloud Infinite Scale: MIGRATION

== Badges need to be provided manually ==

**This project is under heavy development, it's not in a working state yet!**

## Install

You can download prebuilt binaries from the GitHub releases or from our [download mirrors](http://download.owncloud.com/ocis/migration/). For instructions how to install this on your platform you should take a look at our [documentation](https://owncloud.github.io/extensions/ocis_migration/)

## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.13.

```console
git clone https://github.com/owncloud/ocis-migration.git
cd ocis-migration

make generate build

./bin/ocis-migration -h
```

## Security

If you find a security issue please contact security@owncloud.com first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2020 ownCloud GmbH <https://owncloud.com>
```
