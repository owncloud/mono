Enhancement: only send create home request if an account has been migrated

This change adds a check if an account has been migrated by getting it from the
ocis-accounts service. If no account is returned it means it hasn't been migrated.

https://github.com/owncloud/mono/proxy/issues/52
https://github.com/owncloud/mono/proxy/pull/63
