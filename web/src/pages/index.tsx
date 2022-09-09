import Head from "next/head"
import { DocBox, DocBoxSection } from "src/components/DOTs/doc-box"
import { Header } from "src/components/Header"

export default function Home() {
	return (
		<>
			<Head>
				<link rel="icon" href="/favicon.ico" />
				<meta name="description" content="" />
				<title>DOTs</title>
			</Head>
			<Header />
			<main>
				<DocBoxSection title="Getting Started" description="Learn how to get started with the documentation.">
					<DocBox
						title="Introduction"
						description="Learn about the documentation and the code you have purchased."
						href="/docs/introduction"
						done
					/>
					<DocBox
						title="Getting Started"
						description="Learn how to get started."
						href="/docs/getting-started"
					/>
					<DocBox
						title="Reporting Issues"
						description="Learn how to report issues with the documentation or the code or maybe you want to join the Discord server for help."
						href="/docs/getting-in-touch"
					/>
					<DocBox
						title="Writing frontend code"
						description="Learn how to write frontend code."
						href="/docs/writing-frontend-code"
					/>
					<DocBox
						title="Writing backend code"
						description="Learn how to write backend code."
						href="/docs/writing-backend-code"
					/>
					<DocBox
						title="Deploying to production and staging"
						description="Learn how to deploy to production and staging using the CDK."
						href="/docs/deploying"
					/>
					<DocBox
						title="CI/CD"
						description="Learn how to use CI/CD to deploy to production and staging and run your tests"
						href="/docs/ci-cd"
					/>
				</DocBoxSection>
			</main>
		</>
	)
}
