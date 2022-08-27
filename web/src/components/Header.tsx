import Link from "next/link"
import { homeRoute, loginRoute, logoutRoute, registerRoute } from "src/pages/routes"

export function Header() {
	return (
		<header>
			<h1>
				<Link href={homeRoute()}>
					<a>AWSAppSyncGo</a>
				</Link>
			</h1>
			<nav>
				<Link href={loginRoute()}>
					<a>Login</a>
				</Link>
				<Link href={registerRoute()}>
					<a>Register</a>
				</Link>
				<Link href={logoutRoute()}>
					<a>Logout</a>
				</Link>
			</nav>
		</header>
	)
}
