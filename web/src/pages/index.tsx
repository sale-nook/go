import Head from "next/head"
import { Header } from "src/components/Header"
import { useUser } from "src/providers/user"

export default function Home() {
	const { user } = useUser()
	return (
		<>
			<Head>
				<link rel="icon" href="/favicon.ico" />
			</Head>
			<Header />
			<main>
				<div>{user?.id}</div>
			</main>
		</>
	)
}
