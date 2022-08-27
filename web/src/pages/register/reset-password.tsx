import { FormEventHandler, useCallback, useRef, useState } from "react"
import { Auth } from "@aws-amplify/auth"
import Head from "next/head"
import { Header } from "src/components/Header"
import { useRouter } from "next/router"
import { loginRoute } from "src/pages/routes"

export default function ResetPasswordPage() {
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState<string | undefined>()
	const router = useRouter()
	const emailRef = useRef<HTMLInputElement>(null)
	const codeRef = useRef<HTMLInputElement>(null)
	const passwordRef = useRef<HTMLInputElement>(null)
	const confirmPasswordRef = useRef<HTMLInputElement>(null)

	const onSubmit: FormEventHandler<HTMLFormElement> = useCallback(async (e) => {
		e.preventDefault()
		setLoading(true)
		if (!emailRef.current || !passwordRef.current || !codeRef.current || !confirmPasswordRef.current) {
			return
		}

		if (!emailRef.current?.value) {
			setLoading(false)
			router.push(loginRoute())
			return
		}

		if (!codeRef.current.value.match(/^[0-9]{6}$/)) {
			setError("Code must be 6 digits")
			setLoading(false)
			return
		}

		if (passwordRef.current!.value !== confirmPasswordRef.current!.value) {
			setError("Passwords do not match")
			setLoading(false)
			return
		}

		try {
			await Auth.forgotPasswordSubmit(
				decodeURIComponent(emailRef.current!.value),
				codeRef.current!.value,
				passwordRef.current!.value,
			)
			router.push(loginRoute(emailRef.current!.value))
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
					<input type="hidden" name="email" ref={emailRef} required defaultValue={router.query?.email} />
					<label>
						Confirmation Code:
						<input type="text" name="confirmCode" ref={codeRef} required />
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
						Reset Password and go to login
					</button>
				</form>
			</main>
		</>
	)
}
