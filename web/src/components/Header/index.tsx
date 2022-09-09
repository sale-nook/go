import Link from "next/link"
import { loginRoute, logoutRoute, registerRoute } from "src/pages/routes"
import { useUser } from "src/providers/user"
import { Loading } from "src/components/Loading"
import styles from "./Header.module.css"
import { DOTsLogo } from "../DOTs/Logo"

export function Header() {
	const { user, loading } = useUser()

	let nav = null

	if (!user && !loading) {
		nav = (
			<>
				<Link href={loginRoute()}>
					<a className={styles.navLink}>Login</a>
				</Link>
				<Link href={registerRoute()}>
					<a className={styles.navLink}>Register</a>
				</Link>
			</>
		)
	} else if (user && !loading) {
		nav = (
			<Link href={logoutRoute()}>
				<a className={styles.navLink}>Logout</a>
			</Link>
		)
	}

	return (
		<header className={styles.headerContainer}>
			<nav className={styles.navBar}>
				<div className={styles.navColumn}>
					<Link href="/">
						<a className={[styles.navLink, styles.pretendLogo].join(" ")}>
							<DOTsLogo />
						</a>
					</Link>
				</div>
				<div className={styles.navColumn}>
					<Loading loading={loading}>{nav}</Loading>
				</div>
			</nav>
		</header>
	)
}
