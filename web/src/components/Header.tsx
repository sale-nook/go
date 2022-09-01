import Link from "next/link"
import { homeRoute, loginRoute, logoutRoute, registerRoute } from "src/pages/routes"
import { useUser } from "src/providers/user"
import { Loading } from "./Loading"

export function Header() {
	const { user, loading } = useUser()

	let nav = null

	if (!user && !loading) {
		nav = (
			<>
				<Link href={loginRoute()}>
					<a>Login</a>
				</Link>
				<Link href={registerRoute()}>
					<a>Register</a>
				</Link>
			</>
		)
	} else if (user && !loading) {
		nav = (
			<Link href={logoutRoute()}>
				<a>Logout</a>
			</Link>
		)
	}

	return (
		<header>
			<h1>
				<Link href={homeRoute()}>
					<a>AWSAppSyncGo</a>
				</Link>
			</h1>
			<nav>
				<Loading loading={true}>{nav}</Loading>
			</nav>
		</header>
	)
}
