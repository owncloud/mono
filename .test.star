def main(ctx):
  before = []

  stages = [
    demo(ctx),
  ]

  after = [
  ]

  return before + stages + after

def demo(ctx):
  return {
    'kind': 'pipeline',
    'type': 'docker',
    'name': 'demo',
    'platform': {
      'os': 'linux',
      'arch': 'amd64',
    },
    'steps': [
      {
        'name': 'tags-and-master',
        'image': 'owncloud/alpine:latest',
        'pull': 'always',
        'commands': [
          'echo %s' % (ctx.build.ref),
          'echo %s' % (ctx.build.event),
        ],
        'when': {
          'ref': {
            'include': [
              'refs/heads/master',
              'refs/tags/**',
            ],
            'exclude': [
              'refs/tags/**/*'
            ]
          }
        },
      },
      {
        'name': 'submodule-tags',
        'image': 'owncloud/alpine:latest',
        'pull': 'always',
        'commands': [
          'echo %s' % (ctx.build.ref),
          'echo %s' % (ctx.build.event),
        ],
        'when': {
          'ref': [
            'refs/tags/**/v*',
          ]
        },
      },
      {
        'name': 'pulls',
        'image': 'owncloud/alpine:latest',
        'pull': 'always',
        'commands': [
          'echo %s' % (ctx.build.ref),
          'echo %s' % (ctx.build.event),
        ],
        'when': {
          'ref': [
            'refs/pull/**',
          ]
        },
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
