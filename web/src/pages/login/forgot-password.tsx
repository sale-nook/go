import Head from "next/head"
import { Header } from "src/components/Header"
import { FormEventHandler, useCallback, useRef, useState } from "react"
import { Auth } from "@aws-amplify/auth"
import { useRouter } from "next/router"
import { registerRoute, resetPasswordRoute } from "../routes"
import Link from "next/link"

export default function ForgotPasswordPage() {
	const [error, setError] = useState<string | undefined>()
	const [loading, setLoading] = useState(false)
	const emailRef = useRef<HTMLInputElement>(null)
	const router = useRouter()

	const onSubmit: FormEventHandler<HTMLFormElement> = useCallback(async (e) => {
		e.preventDefault()
		setLoading(true)

		if (!emailRef.current) {
			return
		}

		if (!emailRef.current?.value || emailRef.current?.value === "") {
			setError("Email is required")
			setLoading(false)
			return
		}

		try {
			const unsafeEmail = decodeURIComponent(emailRef.current!.value)
			// Send the confirmation code to the user's email.
			await Auth.forgotPassword(unsafeEmail)

			// Send the user to the password reset page.
			router.push(resetPasswordRoute(unsafeEmail))
		} catch (e) {
			console.error(e)
			setError((e as Error).message)
		} finally {
			setLoading(false)
		}
	}, [])

	const errorComponent = error ? <section area-label="form-errors">{error && <p>{error}</p>}</section> : null

	return (
		<>
			<Head>
				<link rel="icon" href="/favicon.ico" />
				<title>Forgot Password</title>
				<meta name="description" content="Forgot Password" />
			</Head>
			<Header />
			<main>
				{errorComponent}
				<form method="post" action="/api/login/forgot" onSubmit={onSubmit}>
					<label>
						Email:
						<input
							type="email"
							name="email"
							ref={emailRef}
							required
							defaultValue={decodeURIComponent(router.query?.email)}
						/>
					</label>
					<button type="submit" disabled={loading}>
						Confirm email and reset password
					</button>
					<Link href={registerRoute()}>
						<a>Register</a>
					</Link>
				</form>
			</main>
		</>
	)
}
