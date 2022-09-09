import Head from "next/head"
import { Header } from "src/components/Header"
import { FormEventHandler, useCallback, useRef, useState } from "react"
import { Auth } from "@aws-amplify/auth"
import { useRouter } from "next/router"
import { forgotPasswordRoute, homeRoute } from "../routes"

export default function LoginPage() {
	const [error, setError] = useState<string | undefined>()
	const [loading, setLoading] = useState(false)
	const emailRef = useRef<HTMLInputElement>(null)
	const passwordRef = useRef<HTMLInputElement>(null)
const router = useRouter()

	const onSubmit: FormEventHandler<HTMLFormElement> = useCallback(async (e) => {
		e.preventDefault()
		setLoading(true)

		if (!emailRef.current || !passwordRef.current) {
			return
		}

		try {
			await Auth.signIn(decodeURIComponent(emailRef.current!.value), passwordRef.current!.value)
			router.push(homeRoute())
		} catch (e) {
			setError((e as Error).message)
		} finally {
			setLoading(false)
		}
	}, [])

	const errorComponent = error ? <section aria-label="form-errors">{error && <p>{error}</p>}</section> : null

	return (
		<>
			<Head>
				<link rel="icon" href="/favicon.ico" />
				<title>Login</title>
				<meta name="description" content="Login to your account" />
			</Head>
			<Header />
			<main>
				{errorComponent}
				<form method="post" action="/api/login" onSubmit={onSubmit}>
					<label>
						Email:
						<input type="email" name="email" ref={emailRef} required defaultValue={router.query?.email} />
					</label>
					<label>
						Password:
						<input type="password" name="password" ref={passwordRef} required />
					</label>
					<button type="submit" disabled={loading}>
						Login
					</button>
					<button
						type="button"
						disabled={loading}
						onClick={() => {
							router.push(forgotPasswordRoute(router.query?.email))
						}}
					>
						I forgot my password
					</button>
				</form>
			</main>
		</>
	)
}
