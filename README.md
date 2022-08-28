# aws-appsync-go
![Backend Build status](https://github.com/davemackintosh/aws-appsync-go/actions/workflows/go.yml/badge.svg?branch=main)
![Frontend Build status](https://github.com/davemackintosh/aws-appsync-go/actions/workflows/web.yml/badge.svg?branch=main)
[![codecov](https://codecov.io/gh/davemackintosh/aws-appsync-go/branch/main/graph/badge.svg?token=milTNQGLWc)](https://codecov.io/gh/davemackintosh/aws-appsync-go)

<blockquote style="text-align: center; text-transform: uppercase">⚠️ This boilerplate requires configuration before use!!! Follow the instructions below carefully. ⚠️</blockquote>
<blockquote>💰 NOTE: Current running costs are unknown, but I will add these as the data becomes more transparent.</blockquote>

# Contents

1. [Description](#description)
2. [Getting Started](#getting-started)
    1. [Installing dependencies](#installing-dependencies)
    2. [Configuring your app](#configuring-your-app)
    3. [Deploying to AWS](#deploying-to-aws)
    4. [Running the frontend](#running-the-frontend)
    5. [Testing your app](#testing-your-app)

----

# Description

This is a boilerplate Golang, NextJS/React application which features the following:

- [Go](https://golang.org/) for all the backend and CDK.
  - Middleware pattern and library for lambda function handlers.
  - Hexagonal architecture/adapter based for the backend.
- [CDK](https://aws.amazon.com/cdk/) (written in Golang) to deploy the application to the cloud.
  - [Lambda serverless functions](https://aws.amazon.com/lambda/)
  - [DynamoDB tables and triggers with point in time recovery](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStarted.Tables.html)
  - [KMS encryption](https://docs.aws.amazon.com/kms/latest/developerguide/services-dynamodb.html)
  - [Cognito User & Identity Pools](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-settings-attributes.html)
  - [AWS AppSync](https://aws.amazon.com/appsync/) authorised and unauthorised access via [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html).
- [Github Actions](https://github.com/features/actions)
  - Runs go tests
  - Runs frontend tests
  - Synthesises & Runs checkov infrastructure as code
  - Builds the frontend
  - Branch based deployment via CDK to AWS (main, staging)
- [Codecov](https://codecov.io/) Code coverage for the backend and frontend
- [Checkov](https://www.checkov.io) for infrastructure as code security checks.
- [CycloneDX](https://github.com/CycloneDX/cyclonedx-gomod) for generating an [SBOM](https://www.cisa.gov/sbom)
- [golangci-lint](https://golangci-lint.run) for static code analysis.
- [Pre-Commit](https://pre-commit.com) for git hook code quality checks.
- [NextJS](https://nextjs.org/) for the frontend.
  - [TypeScript](https://www.typescriptlang.org/)
  - [React](https://reactjs.org/)
  - [ESLint](https://eslint.org/) & [Prettier](https://prettier.io/)
  - [Vitest](https://vitest.dev/) for unit testing.
  - [Cypress](https://www.cypress.io/) for integration testing.
- [Apollo GraphQL](https://www.apollographql.com/)
  - [GraphQL Codegen](https://www.apollographql.com/docs/graphql-tools/codegen/) for generating TypeScript types and Apollo client hooks from our schema automatically.
  - [GraphQL Linting](https://www.apollographql.com/docs/graphql-tools/lint/) for static code analysis.

## Project layout

This repository follows [these standards](https://github.com/golang-standards/project-layout) for project layout.

TL;DR You'll find:

* Private code in `/internal`
* Lambda and other command type functions in `/cmd`
* Frontend in `/web`

----

## Who's this for?

This is a boilerplate for creating a serverless application using the AWS AppSync GraphQL API while using Go for the backend and NextJS for the frontend.

Typical audiences are:

* CTO/CEO looking to save time and money getting a good foundation in place.
* Developers and dev-ops teams looking to get started quickly and cheaply.
* Anyone who wants to get a simple serverless application up and running.

## Getting started

There are a few steps to getting the most out of this boilerplate:

1. [Installing dependencies](#installing-dependencies)
2. [Adding AWS parameter store values](#aws-parameter-store-values)
3. [Configuring your app](#configuring-your-app)
4. [Deploying to AWS](#deploying-to-aws)
5. [Running the frontend](#running-the-frontend)
6. [Testing your app](#testing-your-app)

----

# Installing Dependencies

Before anything, you **must** have a valid AWS account.

All the AWS services are free for the first year, and you can apply for $300 credit [here](https://aws.amazon.com/government-education/sustainability-research-credits/).

## AWS Parameter Store values

To deploy this, the following [parameter store values](https://eu-west-2.console.aws.amazon.com/systems-manager/parameters/?tab=Table) are **required** to be readable by a `cdk` user.
Substitute `{environment}` with any of `staging, production, ci`

```bash
/$APP_NAME/{environment}/GITHUB_ACCESS_TOKEN # Secure String
/$APP_NAME/{environment}/OAUTH_CALLBACK_ROOT # String
```

## System dependencies

* [Golang 1.18+](https://golang.org/)
* [AWS CLI](https://aws.amazon.com/cli/) and [AWS vault](https://github.com/99designs/aws-vault) to setup AWS credentials
* [direnv](https://direnv.net/) to manage environment variables automatically.
* [Pre-Commit](https://pre-commit.com/) for git hook code quality checks to prevent bad commits.
* [nodeJS](https://nodejs.org) & [yarn](https://yarnpkg.com/) for managing/running frontend dependencies and [aws-cdk](https://www.npmjs.com/package/aws-cdk) for the CDK deployments.
* `python3.x`
  * This is needed for serving the graphql schema for graphql-codegen over http. (This will be moved to a NodeJS solution soon but `python -m http.server` is fine for now.)

## Frontend dependencies

The frontend is based on NextJS and React with TypeScript & Vitest and Cypress for testing.

It uses yarn 2.x for managing dependencies.

Running `yarn install` will install all frontend dependencies.

[Read More](./web/README.md) about the frontend.

## Backend dependencies

Go dependencies are managed via [Go modules](https://go.dev/ref/mod) but if you want to download them ahead of time you can run `go mod download` to download them.

----

#### Setting up environmentals.

Edit `./.envrc` and set the following:

```bash
APP_NAME            # The name of the application (affects import paths, should be the git repo name.)
NGROK_DOMAIN        # This is the subdomain on ngrok you'll access the frontend from locally.
AWS_REGION          # This is the region you'll be deploying to.
AWS_ACCOUNT_ID      # This is the account you'll be deploying to.
GITHUB_ACCESS_TOKEN # This is the access token you'll use to access the GitHub API.
```

Go to Github and

1. Generate a new personal access token for the app [here](https://github.com/settings/tokens)
2. Go to repository secrets and add the following:
    ```bash
    ENVIRONMENT
    AWS_REGION
    AWS_ACCOUNT_ID
    GITHUB_ACCESS_TOKEN
    OAUTH_CALLBACK_ROOT
    CODECOV_TOKEN
    ```

---

## Configuring your app

You should have the dependencies' installation instructions AND the environmentals instructions completed above to continue, you should be able to run `init-project` which will rename this project to match your Git repository name.

## Deployment

The CDK supports creating stacks per any environment, supported environments are `production, staging, ci`

```
deploy
```

Or you can deploy a specific environment with

```bash
ENVIRONMENT={production, staging, ci} deploy
```
