# aws-appsync-go
![Backend Build status](https://github.com/davemackintosh/aws-appsync-go/actions/workflows/go.yml/badge.svg?branch=main)
![Frontend Build status](https://github.com/davemackintosh/aws-appsync-go/actions/workflows/web.yml/badge.svg?branch=main)
[![codecov](https://codecov.io/gh/davemackintosh/aws-appsync-go/branch/main/graph/badge.svg?token=milTNQGLWc)](https://codecov.io/gh/davemackintosh/aws-appsync-go)

This repository follows [these standards](https://github.com/golang-standards/project-layout) for project layout.

TL;DR You'll find:

* GraphQL Schema in `/api`
* Private code in `/internal`
* Lambda and other command type functions in `/cmd`
* Frontend in `/web`

----

## Deployment

The CDK supports creating stacks per any environment, supported environments are `production, staging, ci`

```
cdk deploy --all --outputs-file=./config/apis.json
```

## Dependencies

Install all dependencies using `yarn && go mod download`.

### AWS parameter store

To deploy this to a new stack, the following SSM parameters are required to be readable by the `cdk` user. Substitute `{environment}` with any of `staging, production, ci`

```bash
/{environment}/OAUTH_CALLBACK_ROOT # String
```
