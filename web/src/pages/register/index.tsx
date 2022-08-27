import { FormEventHandler, useCallback, useRef, useState } from "react"
import { Auth } from "@aws-amplify/auth"
import Head from "next/head"
import { Header } from "src/components/Header"
import { useRouter } from "next/router"
import { registerConfirmRoute } from "src/pages/routes"

export default function RegisterPage() {
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState<string | undefined>()
	const router = useRouter()
	const emailRef = useRef<HTMLInputElement>(null)
	const passwordRef = useRef<HTMLInputElement>(null)
	const confirmPasswordRef = useRef<HTMLInputElement>(null)

	const onSubmit: FormEventHandler<HTMLFormElement> = useCallback(async (e) => {
		e.preventDefault()
		setLoading(true)
		if (!emailRef.current || !passwordRef.current) {
			return
		}

		if (passwordRef.current!.value !== confirmPasswordRef.current!.value) {
			setError("Passwords do not match")
			setLoading(false)
			return
		}

		try {
			const res = await Auth.signUp({
				username: emailRef.current!.value,
				password: passwordRef.current!.value,
				attributes: {
					email: emailRef.current!.value,
				},
				autoSignIn: {
					enabled: true,
				},
			})

			console.log("RES", res)

			router.push(registerConfirmRoute(emailRef.current!.value))
		} catch (e) {
			setError((e as Error).message)
			setLoading(false)
		} finally {
			setLoading(false)
		}
	}, [])

	const errorComponent = error ? <section area-label="form-errors">{error && <p>{error}</p>}</section> : null

	return (
		<>
			<Head>
				<link rel="icon" href="/favicon.ico" />
			</Head>
			<Header />
			<main>
				{errorComponent}
				<form method="post" action="/api/register" onSubmit={onSubmit}>
					<label>
						Email:
						<input type="email" name="email" ref={emailRef} required />
					</label>
					<label>
						Password:
						<input type="password" name="password" ref={passwordRef} required />
					</label>
					<label>
						Confirm password:
						<input type="password" name="confirm-password" ref={confirmPasswordRef} required />
					</label>
					<button type="submit" disabled={loading}>
						Register
					</button>
				</form>
			</main>
		</>
	)
}
