config = {
  'modules': {
    'accounts': 'frontend',
    'glauth':'',
    'reva':'',
    'ocis-pkg':'',
    'phoenix':'',
    'store':'',
    'ocs':'',
    'webdav':'',
    'thumbnails':'',
    'proxy':'',
    'settings':'frontend',
    'konnectd':'',
  },
  'apiTests': {
    'coreBranch': 'master',
    'coreCommit': '39bcf3874f2e388bbe7a7802e2f1e3c6a4b8216a',
    'numberOfParts': 2
  },
  'uiTests': {
    'phoenixBranch': 'master',
    'phoenixCommit': '2ccec3966abebc7cff54cfdbd0aa9942253584a7',
    'suites': {
      'phoenixWebUI1': [
        'webUICreateFilesFolders',
        'webUIDeleteFilesFolders',
        'webUIFavorites',
        'webUIFiles',
        'webUILogin',
        'webUINotifications',
      ],
      'phoenixWebUI2': [
        'webUIPrivateLinks',
        'webUIRenameFiles',
        'webUIRenameFolders',
        'webUITrashbin',
        'webUIUpload',
        'webUIAccount',
        # All tests in the following suites are skipped currently
        # so they won't run now but when they are enabled they will run
        'webUIRestrictSharing',
        'webUISharingAutocompletion',
        'webUISharingInternalGroups',
        'webUISharingInternalUsers',
        'webUISharingPermissionsUsers',
        'webUISharingFilePermissionsGroups',
        'webUISharingFolderPermissionsGroups',
        'webUISharingFolderAdvancedPermissionsGroups',
        'webUIResharing',
        'webUISharingPublic',
        'webUISharingPublicDifferentRoles',
        'webUISharingAcceptShares',
        'webUISharingFilePermissionMultipleUsers',
        'webUISharingFolderPermissionMultipleUsers',
        'webUISharingFolderAdvancedPermissionMultipleUsers',
        'webUISharingNotifications',
      ],
    }
  }
}

def getUITestSuiteNames():
  return config['uiTests']['suites'].keys()

def getUITestSuites():
  return config['uiTests']['suites']

def getCoreApiTestPipelineNames():
  names = []
  for runPart in range(1, config['apiTests']['numberOfParts'] + 1):
    names.append('coreApiTests-%s' % runPart)
  return names

def main(ctx):
  before = [
  ]

  for module in config['modules']:
    before.append(linting(ctx, module))
  before += testPipelines(ctx)
  
  stages = [
  ]

  for module in config['modules']:
   stages.append(unitTests(ctx, module))

  after = [
  ]

  return before + stages + after

def testPipelines(ctx):

  pipelines = [
    localApiTestsOcStorage(ctx, config['apiTests']['coreBranch'], config['apiTests']['coreCommit'])
  ]

  for runPart in range(1, config['apiTests']['numberOfParts'] + 1):
    pipelines.append(coreApiTests(ctx, config['apiTests']['coreBranch'], config['apiTests']['coreCommit'], runPart, config['apiTests']['numberOfParts']))

  pipelines += uiTests(ctx, config['uiTests']['phoenixBranch'], config['uiTests']['phoenixCommit'])
  return pipelines

def linting(ctx, module):
  steps = generate(module) + [
    {
      'name': 'vet',
      'image': 'webhippie/golang:1.13',
      'pull': 'always',
      'commands': [
        'cd %s' % (module),
        'make vet',
      ],
      'volumes': [
        {
          'name': 'gopath',
          'path': '/srv/app',
        },
      ],
    },
    {
      'name': 'staticcheck',
      'image': 'webhippie/golang:1.13',
      'pull': 'always',
      'commands': [
        'cd %s' % (module),
        'make staticcheck',
      ],
      'volumes': [
        {
          'name': 'gopath',
          'path': '/srv/app',
        },
      ],
    },
    {
      'name': 'lint',
      'image': 'webhippie/golang:1.13',
      'pull': 'always',
      'commands': [
        'cd %s' % (module),
        'make lint',
      ],
      'volumes': [
        {
          'name': 'gopath',
          'path': '/srv/app',
        },
      ],
    },
  ]

  if config['modules'][module] == 'frontend': 
    steps = frontend(module) + steps

  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'linting-%s' % (module),
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': steps,
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**',
      ],
    },
  }

def unitTests(ctx, module):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'unitTests-%s' % (module),
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps':
    [
      {
        'name': 'test',
        'image': 'webhippie/golang:1.13',
        'pull': 'always',
        'commands': [
          'cd %s' % (module),
          'make test',
        ],
        'volumes': [
          {
            'name': 'gopath',
            'path': '/srv/app',
          },
        ],
      },
      {
        'name': 'codacy',
        'image': 'plugins/codacy:1',
        'pull': 'always',
        'settings': {
          'token': {
            'from_secret': 'codacy_token',
          },
        },
      },
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**',
      ],
    },
  }

def localApiTestsOcStorage(ctx, coreBranch = 'master', coreCommit = ''):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'localApiTestsOcStorage',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps':
      generate('ocis') +
      build() +
      ocisServer() +
      cloneCoreRepos(coreBranch, coreCommit) + [
      {
        'name': 'localApiTestsOcStorage',
        'image': 'owncloudci/php:7.2',
        'pull': 'always',
        'environment' : {
          'TEST_SERVER_URL': 'https://ocis-server:9200',
          'OCIS_REVA_DATA_ROOT': '/srv/app/tmp/reva/',
          'SKELETON_DIR': '/srv/app/tmp/testing/data/apiSkeleton',
          'TEST_OCIS':'true',
          'BEHAT_FILTER_TAGS': '~@skipOnOcis-OC-Storage',
          'PATH_TO_CORE': '/srv/app/testrunner'
        },
        'commands': [
          'cd ocis',
          'make test-acceptance-api',
        ],
        'volumes': [{
          'name': 'gopath',
          'path': '/srv/app',
        }]
      },
    ],
    'services':
      redis(),
    'volumes': [
      {
        'name': 'gopath',
        'temp': {},
      },
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**',
      ],
    },
  }

def coreApiTests(ctx, coreBranch = 'master', coreCommit = '', part_number = 1, number_of_parts = 1):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'coreApiTests-%s' % (part_number),
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps':
      generate('ocis') +
      build() +
      ocisServer() +
      cloneCoreRepos(coreBranch, coreCommit) + [
      {
        'name': 'oC10ApiTests-%s' % (part_number),
        'image': 'owncloudci/php:7.2',
        'pull': 'always',
        'environment' : {
          'TEST_SERVER_URL': 'https://ocis-server:9200',
          'OCIS_REVA_DATA_ROOT': '/srv/app/tmp/reva/',
          'SKELETON_DIR': '/srv/app/tmp/testing/data/apiSkeleton',
          'TEST_OCIS':'true',
          'BEHAT_FILTER_TAGS': '~@notToImplementOnOCIS&&~@toImplementOnOCIS&&~comments-app-required&&~@federation-app-required&&~@notifications-app-required&&~systemtags-app-required&&~@local_storage',
          'DIVIDE_INTO_NUM_PARTS': number_of_parts,
          'RUN_PART': part_number,
          'EXPECTED_FAILURES_FILE': '/drone/src/ocis/tests/acceptance/expected-failures-on-OC-storage.txt'
        },
        'commands': [
          'cd /srv/app/testrunner',
          'make test-acceptance-api',
        ],
        'volumes': [{
          'name': 'gopath',
          'path': '/srv/app',
        }]
      },
    ],
    'services':
      redis(),
    'volumes': [
      {
        'name': 'gopath',
        'temp': {},
      },
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**',
      ],
    },
  }

def uiTests(ctx, phoenixBranch, phoenixCommit):
  suiteNames = getUITestSuiteNames()
  return [uiTestPipeline(suiteName, phoenixBranch, phoenixCommit) for suiteName in suiteNames]

def uiTestPipeline(suiteName, phoenixBranch = 'master', phoenixCommit = ''):
  suites = getUITestSuites()
  paths = ""
  for path in suites[suiteName]:
    paths = paths + "tests/acceptance/features/" + path + " "

  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': suiteName,
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps':
      generate('ocis') +
      build() +
      ocisServer() + [
      {
        'name': 'webUITests',
        'image': 'owncloudci/nodejs:11',
        'pull': 'always',
        'environment': {
          'SERVER_HOST': 'https://ocis-server:9200',
          'BACKEND_HOST': 'https://ocis-server:9200',
          'RUN_ON_OCIS': 'true',
          'OCIS_REVA_DATA_ROOT': '/srv/app/tmp/reva',
          'OCIS_SKELETON_DIR': '/srv/app/testing/data/webUISkeleton',
          'PHOENIX_CONFIG': '/drone/src/ocis/tests/config/drone/ocis-config.json',
          'TEST_TAGS': 'not @skipOnOCIS and not @skip',
          'LOCAL_UPLOAD_DIR': '/uploads',
          'NODE_TLS_REJECT_UNAUTHORIZED': 0,
          'TEST_PATHS': paths,
        },
        'commands': [
          'git clone -b master --depth=1 https://github.com/owncloud/testing.git /srv/app/testing',
          'git clone -b %s --single-branch --no-tags https://github.com/owncloud/phoenix.git /srv/app/phoenix' % (phoenixBranch),
          'cp -r /srv/app/phoenix/tests/acceptance/filesForUpload/* /uploads',
          'cd /srv/app/phoenix',
        ] + ([
          'git checkout %s' % (phoenixCommit)
        ] if phoenixCommit != '' else []) + [
          'yarn install-all',
          'yarn run acceptance-tests-drone'
        ],
        'volumes': [{
          'name': 'gopath',
          'path': '/srv/app',
        },
        {
          'name': 'uploads',
          'path': '/uploads'
        }]
      },
    ],
    'services':
      redis() +
      selenium(),
    'volumes': [
      {
        'name': 'gopath',
        'temp': {},
      },
      {
        'name': 'uploads',
        'temp': {}
      }
    ],
    'trigger': {
      'ref': [
        'refs/heads/master',
        'refs/tags/**',
        'refs/pull/**',
      ],
    },
  }

def generate(module):
  return [
    {
      'name': 'generate',
      'image': 'webhippie/golang:1.13',
      'pull': 'always',
      'commands': [
        'cd %s' % (module),
        'make generate',
      ],
      'volumes': [
        {
          'name': 'gopath',
          'path': '/srv/app',
        },
      ],
    }
  ]

def frontend(module):
  return [
    {
      'name': 'frontend',
      'image': 'webhippie/nodejs:latest',
      'pull': 'always',
      'commands': [
        'cd %s' % (module),
        'yarn install --frozen-lockfile',
        'yarn lint',
        'yarn test',
        'yarn build',
      ],
    }
  ]

def ocisServer():
  return [
    {
      'name': 'ocis-server',
      'image': 'webhippie/golang:1.13',
      'pull': 'always',
      'detach': True,
      'environment' : {
        'OCIS_LOG_LEVEL': 'debug',

        'REVA_STORAGE_HOME_DATA_TEMP_FOLDER': '/srv/app/tmp/',
        'REVA_STORAGE_LOCAL_ROOT': '/srv/app/tmp/reva/root',
        'REVA_STORAGE_OWNCLOUD_DATADIR': '/srv/app/tmp/reva/data',
        'REVA_STORAGE_OC_DATA_TEMP_FOLDER': '/srv/app/tmp/',
        'REVA_STORAGE_OWNCLOUD_REDIS_ADDR': 'redis:6379',
        'REVA_LDAP_IDP': 'https://ocis-server:9200',
        'REVA_OIDC_ISSUER': 'https://ocis-server:9200',
        'PROXY_OIDC_ISSUER': 'https://ocis-server:9200',
        'REVA_STORAGE_OC_DATA_SERVER_URL': 'http://ocis-server:9164/data',
        'REVA_DATAGATEWAY_URL': 'https://ocis-server:9200/data',
        'REVA_FRONTEND_URL': 'https://ocis-server:9200',
        'PHOENIX_WEB_CONFIG': '/drone/src/ocis/tests/config/drone/ocis-config.json',
        'KONNECTD_IDENTIFIER_REGISTRATION_CONF': '/drone/src/ocis/tests/config/drone/identifier-registration.yml',
        'KONNECTD_ISS': 'https://ocis-server:9200',
        'KONNECTD_TLS': 'true',
      },
      'commands': [
        'apk add mailcap', # install /etc/mime.types
        'mkdir -p /srv/app/tmp/reva',
        'ocis/bin/ocis server'
      ],
      'volumes': [
        {
          'name': 'gopath',
          'path': '/srv/app'
        },
      ]
    },
  ]

def cloneCoreRepos(coreBranch, coreCommit):
  return [
    {
      'name': 'clone-core-repos',
      'image': 'owncloudci/php:7.2',
      'pull': 'always',
      'commands': [
        'git clone -b master --depth=1 https://github.com/owncloud/testing.git /srv/app/tmp/testing',
        'git clone -b %s --single-branch --no-tags https://github.com/owncloud/core.git /srv/app/testrunner' % (coreBranch),
        'cd /srv/app/testrunner',
      ] + ([
        'git checkout %s' % (coreCommit)
      ] if coreCommit != '' else []),
      'volumes': [{
        'name': 'gopath',
        'path': '/srv/app',
      }]
    }
  ]

def build():
  return [
    {
      'name': 'build',
      'image': 'webhippie/golang:1.13',
      'pull': 'always',
      'commands': [
        'cd ocis',
        'make build',
      ],
      'volumes': [
        {
          'name': 'gopath',
          'path': '/srv/app',
        },
      ],
    },
  ]

def redis():
  return [
    {
      'name': 'redis',
      'image': 'webhippie/redis',
      'pull': 'always',
      'environment': {
        'REDIS_DATABASES': 1
      },
    }
  ]

def selenium():
  return [
    {
      'name': 'selenium',
      'image': 'selenium/standalone-chrome-debug:3.141.59-20200326',
      'pull': 'always',
      'volumes': [{
          'name': 'uploads',
          'path': '/uploads'
      }],
    }
  ]
