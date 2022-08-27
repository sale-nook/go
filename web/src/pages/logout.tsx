import { Auth } from "@aws-amplify/auth"
import Link from "next/link"
import { useRouter } from "next/router"
import { useEffect } from "react"
import { useUser } from "src/providers/user"
import { homeRoute } from "./routes"

export default function LogoutPage() {
	const router = useRouter()
	const { setUser } = useUser()
	useEffect(() => {
		;(async () => {
			await Auth.signOut()
			setUser(null)
			router.push(homeRoute())
		})()
	})
	return (
		<div>
			<h1>Logout</h1>
			<p>You have been sucurely logged out.</p>
			<Link href={homeRoute()}>
				<a>Go to home</a>
			</Link>
		</div>
	)
}
