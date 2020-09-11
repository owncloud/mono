# Changelog for [unreleased] (UNRELEASED)

The following sections list the changes in ocis-settings unreleased.

[unreleased]: https://github.com/owncloud/ocis-settings/compare/v0.3.1...master

## Summary

* Bugfix - Fix loading and saving system scoped values: [#66](https://github.com/owncloud/ocis-settings/pull/66)
* Bugfix - Complete input validation: [#66](https://github.com/owncloud/ocis-settings/pull/66)
* Change - Add filter option for bundle ids in ListBundles and ListRoles: [#59](https://github.com/owncloud/ocis-settings/pull/59)
* Change - Reuse roleIDs from the metadata context: [#69](https://github.com/owncloud/ocis-settings/pull/69)
* Change - Update ocis-pkg/v2: [#72](https://github.com/owncloud/ocis-settings/pull/72)

## Details

* Bugfix - Fix loading and saving system scoped values: [#66](https://github.com/owncloud/ocis-settings/pull/66)

   We fixed loading and saving system scoped values. Those are now saved without an account uuid,
   so that the value can be loaded by other accounts as well.

   https://github.com/owncloud/ocis-settings/pull/66


* Bugfix - Complete input validation: [#66](https://github.com/owncloud/ocis-settings/pull/66)

   There was one handler function without input validation. We implemented the input validation
   for `ValueService.ReadValueByUniqueIdentifiers`.

   https://github.com/owncloud/ocis-settings/pull/66


* Change - Add filter option for bundle ids in ListBundles and ListRoles: [#59](https://github.com/owncloud/ocis-settings/pull/59)

   We added bundle ids as filter option for ListBundles and ListRoles and a new endpoint for
   fetching a permission by ID.

   https://github.com/owncloud/ocis-settings/pull/59


* Change - Reuse roleIDs from the metadata context: [#69](https://github.com/owncloud/ocis-settings/pull/69)

   The roleIDs of the authenticated user are coming in from the metadata context. Since we decided
   to move the role assignments over to the accounts service we need to start trusting those
   roleIDs from the metadata context instead of reloading them from disk on each request.

   https://github.com/owncloud/ocis-settings/pull/69


* Change - Update ocis-pkg/v2: [#72](https://github.com/owncloud/ocis-settings/pull/72)

   The helper we used from ocis-pkg for extracting roleIDs from the metadata context has moved to
   another package, so we needed to update as well.

   https://github.com/owncloud/ocis-settings/pull/72
   https://github.com/owncloud/ocis-pkg/pull/60

# Changelog for [0.3.1] (2020-08-27)

The following sections list the changes in ocis-settings 0.3.1.

[0.3.1]: https://github.com/owncloud/ocis-settings/compare/v0.3.0...v0.3.1

## Summary

* Bugfix - Fix fetching bundles in settings UI: [#61](https://github.com/owncloud/ocis-settings/pull/61)

## Details

* Bugfix - Fix fetching bundles in settings UI: [#61](https://github.com/owncloud/ocis-settings/pull/61)

   We fixed the settings UI to use the changed API endpoint `BundleService.ListBundles`
   properly.

   https://github.com/owncloud/ocis-settings/pull/61

# Changelog for [0.3.0] (2020-08-26)

The following sections list the changes in ocis-settings 0.3.0.

[0.3.0]: https://github.com/owncloud/ocis-settings/compare/v0.2.0...v0.3.0

## Summary

* Change - Filter settings by permissions: [#99](https://github.com/owncloud/product/issues/99)

## Details

* Change - Filter settings by permissions: [#99](https://github.com/owncloud/product/issues/99)

   `BundleService.GetBundle` and `BundleService.ListBundles` are now filtered by READ
   permissions in the role of the authenticated user. This prevents settings from being visible
   to the user when their role doesn't have appropriate permissions.

   https://github.com/owncloud/product/issues/99
   https://github.com/owncloud/ocis-settings/pull/48

# Changelog for [0.2.0] (2020-08-20)

The following sections list the changes in ocis-settings 0.2.0.

[0.2.0]: https://github.com/owncloud/ocis-settings/compare/v0.1.0...v0.2.0

## Summary

* Change - Add role service: [#110](https://github.com/owncloud/product/issues/110)
* Change - Rename endpoints and message types: [#36](https://github.com/owncloud/ocis-settings/issues/36)
* Change - Use UUIDs instead of alphanumeric identifiers: [#46](https://github.com/owncloud/ocis-settings/pull/46)

## Details

* Change - Add role service: [#110](https://github.com/owncloud/product/issues/110)

   We added service endpoints for registering roles and maintaining permissions.

   https://github.com/owncloud/product/issues/110
   https://github.com/owncloud/ocis-settings/issues/10
   https://github.com/owncloud/ocis-settings/pull/47


* Change - Rename endpoints and message types: [#36](https://github.com/owncloud/ocis-settings/issues/36)

   We decided to rename endpoints and message types to be less verbose. Specifically,
   `SettingsBundle` became `Bundle`, `Setting` (inside a bundle) kept its name and
   `SettingsValue` became `Value`.

   https://github.com/owncloud/ocis-settings/issues/36
   https://github.com/owncloud/ocis-settings/issues/32
   https://github.com/owncloud/ocis-settings/pull/46


* Change - Use UUIDs instead of alphanumeric identifiers: [#46](https://github.com/owncloud/ocis-settings/pull/46)

   `Bundles`, `Settings` and `Values` were identified by a set of alphanumeric identifiers so
   far. We switched to UUIDs in order to achieve a flat file hierarchy on disk. Referencing the
   respective entities by their alphanumeric identifiers (as used in UI code) is still
   supported.

   https://github.com/owncloud/ocis-settings/pull/46

# Changelog for [0.1.0] (2020-08-17)

The following sections list the changes in ocis-settings 0.1.0.

[0.1.0]: https://github.com/owncloud/ocis-settings/compare/6fdbbd7e8eb39f18ada1a8a3c45a1c925e239889...v0.1.0

## Summary

* Bugfix - Adjust UUID validation to be more tolerant: [#41](https://github.com/owncloud/ocis-settings/issues/41)
* Bugfix - Fix runtime error when type asserting on nil value: [#38](https://github.com/owncloud/ocis-settings/pull/38)
* Bugfix - Fix multiple submits on string and number form elements: [#745](https://github.com/owncloud/owncloud-design-system/issues/745)
* Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#39](https://github.com/owncloud/ocis-settings/pull/39)
* Change - Dynamically add navItems for extensions with settings bundles: [#25](https://github.com/owncloud/ocis-settings/pull/25)
* Change - Introduce input validation: [#22](https://github.com/owncloud/ocis-settings/pull/22)
* Change - Use account uuid from x-access-token: [#14](https://github.com/owncloud/ocis-settings/pull/14)
* Change - Use server config variable from ocis-web: [#34](https://github.com/owncloud/ocis-settings/pull/34)
* Enhancement - Remove paths from Makefile: [#33](https://github.com/owncloud/ocis-settings/pull/33)
* Enhancement - Extend the docs: [#11](https://github.com/owncloud/ocis-settings/issues/11)
* Enhancement - Update ocis-pkg/v2: [#42](https://github.com/owncloud/ocis-settings/pull/42)

## Details

* Bugfix - Adjust UUID validation to be more tolerant: [#41](https://github.com/owncloud/ocis-settings/issues/41)

   The UUID now allows any alphanumeric character and "-", "_", ".", "+" and "@" which can also
   allow regular user names.

   https://github.com/owncloud/ocis-settings/issues/41


* Bugfix - Fix runtime error when type asserting on nil value: [#38](https://github.com/owncloud/ocis-settings/pull/38)

   Fixed the case where an account UUID present in the context is nil, and type asserting it as a
   string would produce a runtime error.

   https://github.com/owncloud/ocis-settings/issues/37
   https://github.com/owncloud/ocis-settings/pull/38


* Bugfix - Fix multiple submits on string and number form elements: [#745](https://github.com/owncloud/owncloud-design-system/issues/745)

   We had a bug with keyboard event listeners triggering multiple submits on input fields. This
   was recently fixed in the ownCloud design system (ODS). We rolled out that bugfix to the
   settings ui as well.

   https://github.com/owncloud/owncloud-design-system/issues/745
   https://github.com/owncloud/owncloud-design-system/pull/768
   https://github.com/owncloud/ocis-settings/pulls/31


* Bugfix - Build docker images with alpine:latest instead of alpine:edge: [#39](https://github.com/owncloud/ocis-settings/pull/39)

   ARM builds were failing when built on alpine:edge, so we switched to alpine:latest instead.

   https://github.com/owncloud/ocis-settings/pull/39


* Change - Dynamically add navItems for extensions with settings bundles: [#25](https://github.com/owncloud/ocis-settings/pull/25)

   We now make use of a new feature in ocis-web-core, allowing us to add navItems not only through
   configuration, but also after app initialization. With this we now have navItems available
   for all extensions within the settings ui, that have at least one settings bundle registered.

   https://github.com/owncloud/ocis-settings/pull/25


* Change - Introduce input validation: [#22](https://github.com/owncloud/ocis-settings/pull/22)

   We set up input validation, starting with enforcing alphanumeric identifier values and UUID
   format on account uuids. As a result, traversal into parent folders is not possible anymore. We
   also made sure that get and list requests are side effect free, i.e. not creating any folders.

   https://github.com/owncloud/ocis-settings/issues/15
   https://github.com/owncloud/ocis-settings/issues/16
   https://github.com/owncloud/ocis-settings/issues/19
   https://github.com/owncloud/ocis-settings/pull/22


* Change - Use account uuid from x-access-token: [#14](https://github.com/owncloud/ocis-settings/pull/14)

   We are now using an ocis-pkg middleware for extracting the account uuid of the authenticated
   user from the `x-access-token` of the http request header and inject it into the Identifier
   protobuf messages wherever possible. This allows us to use `me` instead of an actual account
   uuid, when the request comes through the proxy.

   https://github.com/owncloud/ocis-settings/pull/14


* Change - Use server config variable from ocis-web: [#34](https://github.com/owncloud/ocis-settings/pull/34)

   We are not providing an api url anymore but use the server url from ocis-web config instead. This
   still - as before - requires that ocis-proxy is in place for routing requests to ocis-settings.

   https://github.com/owncloud/ocis-settings/pull/34


* Enhancement - Remove paths from Makefile: [#33](https://github.com/owncloud/ocis-settings/pull/33)

   We have a variable for the proto files path in our Makefile, but were not using it. Changed the
   Makefile to use the PROTO_SRC variable where possible.

   https://github.com/owncloud/ocis-settings/pull/33


* Enhancement - Extend the docs: [#11](https://github.com/owncloud/ocis-settings/issues/11)

   We have extended the documentation by adding a chapter about settings values.

   https://github.com/owncloud/ocis-settings/issues/11
   https://github.com/owncloud/ocis-settings/pulls/28


* Enhancement - Update ocis-pkg/v2: [#42](https://github.com/owncloud/ocis-settings/pull/42)

   Update to ocis-pkg/v2 v2.2.2-0.20200812103920-db41b5a3d14d

   https://github.com/owncloud/ocis-settings/pull/42

