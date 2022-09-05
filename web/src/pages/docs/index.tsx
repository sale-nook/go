import Head from "next/head"
import Link from "next/link"
import { Header } from "src/components/Header"
import { useUser } from "src/providers/user"
import { Contents } from "src/components/docs/Contents"

export default function DocsIndex() {
	return (
		<>
			<Head>
				<link rel="icon" href="/favicon.ico" />
			</Head>
			<Header />
			<main>
				<h1>Docs</h1>

				<section id="description">
					<p>This is a boilerplate Golang, NextJS/React application which features the following:</p>

					<ul>
						<li>
							<Link href="https://golang.org/">
								<a>GO</a>
							</Link>{" "}
							for all the backend and CDK.
							<ul>
								<li>Middleware pattern and library for lambda function handlers.</li>
								<li>Hexagonal architecture/adapter based for the backend.</li>
							</ul>
						</li>
						<li>
							<Link href="https://aws.amazon.com/cdk/">
								<a>CDK</a>
							</Link>{" "}
							(written in Golang) to deploy the application to the cloud.
							<ul>
								<li>
									<Link href="https://aws.amazon.com/lambda/">
										<a>Lambda serverless functions</a>
									</Link>
								</li>
								<li>
									<Link href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStarted.Tables.html">
										<a>DynamoDB tables and triggers with point in time recovery</a>
									</Link>
								</li>
								<li>
									<Link href="https://docs.aws.amazon.com/kms/latest/developerguide/services-dynamodb.html">
										<a>KMS encryption</a>
									</Link>
								</li>
								<li>
									<Link href="https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-settings-attributes.html">
										<a>Cognito User & Identity Pools</a>
									</Link>
								</li>
								<li>
									<Link href="https://aws.amazon.com/appsync/">
										<a>AWS AppSync</a>
									</Link>{" "}
									authorised and unauthorised access via{" "}
									<Link href="https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html">
										<a>IAM roles.</a>
									</Link>
								</li>
							</ul>
						</li>
						<li>
							<Link href="https://github.com/features/actions">
								<a>Github Actions</a>
							</Link>
							<ul>
								<li>Runs go tests</li>
								<li>Runs frontend tests</li>
								<li>Synthesises &amp; Runs checkov infrastructure as code</li>
								<li>Builds the frontend</li>
								<li>Branch based deployment via CDK to AWS (main, staging)</li>
							</ul>
						</li>
						<li>
							<Link href="https://codecov.io/">
								<a>Codecov</a>
							</Link>{" "}
							Code coverage for the backend and frontend
						</li>
						<li>
							<Link href="https://www.checkov.io">
								<a>Checkov</a>
							</Link>{" "}
							for infrastructure as code security checks.
						</li>
						<li>
							<Link href="https://github.com/CycloneDX/cyclonedx-gomod">
								<a>CycloneDX</a>
							</Link>{" "}
							for generating an{" "}
							<Link href="https://www.cisa.gov/sbom">
								<a>SBOM</a>
							</Link>
						</li>
						<li>
							<Link href="https://golangci-lint.run">
								<a>golangci-lint</a>
							</Link>{" "}
							for static code analysis.
						</li>
						<li>
							<Link href="https://pre-commit.com">
								<a>Pre-Commit</a>
							</Link>{" "}
							for git hook code quality checks.
						</li>
						<li>
							<Link href="https://nextjs.org/">
								<a>NextJS</a>
							</Link>{" "}
							for the frontend.
							<ul>
								<li>
									<Link href="https://www.typescriptlang.org/">
										<a>TypeScript</a>
									</Link>
								</li>
								<li>
									<Link href="https://reactjs.org/">
										<a>React</a>
									</Link>
								</li>
								<li>
									<Link href="https://eslint.org">
										<a>ESLint</a>
									</Link>{" "}
									&amp{" "}
									<Link href="https://prettier.io/">
										<a>Prettier</a>
									</Link>
								</li>
								<li>
									<Link href="https://vitest.dev">
										<a>Vitest</a>
									</Link>{" "}
									for unit testing.
								</li>
								<li>
									<Link href="https://www.cypress.io">
										<a>Cypress</a>
									</Link>{" "}
									for integration testing.
								</li>
								<li>
									<Link href="https://www.apollographql.com">
										<a>Apollo GraphQL</a>
									</Link>
								</li>
								<li>
									<Link href="https://www.apollographql.com/docs/graphql-tools/codegen">
										<a>GraphQL Codegen</a>
									</Link>{" "}
									for generating TypeScript types and Apollo client hooks from our schema
									automatically.
								</li>
								<li>
									<Link href="https://www.apollographql.com/docs/graphql-tools/lint">
										<a>GraphQL Linting</a>
									</Link>{" "}
									for static code analysis.
								</li>
							</ul>
						</li>
					</ul>
				</section>

				<section id="contents">
					<h2>Contents</h2>
					<Contents />
				</section>
			</main>
		</>
	)
}
