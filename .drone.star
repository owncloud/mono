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
  }
}
def main(ctx):
  before = [
  ]

  for module in config['modules']:
    before.append(linting(ctx, module))

  stages = [
  ]

  for module in config['modules']:
   stages.append(unitTests(ctx, module))

  after = [
  ]

  return before + stages + after

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