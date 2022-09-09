import Link from "next/link"

export function Contents() {
	return (
		<ol>
			<li>
				<Link href="/docs/getting-started">Getting Started</Link>
				<ol>
					<li>
						<Link href="/docs/getting-started/installing-dependencies">
							<a>Installing dependencies</a>
						</Link>
					</li>
					<li>
						<Link href="/docs/getting-started/configuring-your-app">
							<a>Configuring your app</a>
						</Link>
					</li>
					<li>
						<Link href="/docs/getting-started/deploying-to-aws">
							<a>Deploying to AWS</a>
						</Link>
					</li>
					<li>
						<Link href="/docs/getting-started/running-the-frontend">
							<a>Running the frontend</a>
						</Link>
					</li>
					<li>
						<Link href="/docs/getting-started/testing-your-new-app">
							<a>Testing your new app</a>
						</Link>
					</li>
				</ol>
			</li>
		</ol>
	)
}
