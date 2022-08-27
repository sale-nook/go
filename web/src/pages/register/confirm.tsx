import Head from "next/head"
import { Header } from "src/components/Header"
import { FormEventHandler, useCallback, useRef, useState } from "react"
import { useRouter } from "next/router"
import { Auth } from "@aws-amplify/auth"
import { homeRoute } from "../routes"

export default function ConfirmEmailPage() {
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState<string | undefined>()
	const emailRef = useRef<HTMLInputElement>(null)
	const codeRef = useRef<HTMLInputElement>(null)
	const router = useRouter()

	const onSubmit: FormEventHandler<HTMLFormElement> = useCallback(async (e) => {
		e.preventDefault()
		setLoading(true)
		if (!emailRef.current || !codeRef.current) {
			return
		}

		try {
			await Auth.confirmSignUp(decodeURIComponent(emailRef.current.value), codeRef.current.value)

			router.push(homeRoute())
		} catch (e) {
			setError((e as Error).message)
			setLoading(false)
			return
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
				<form method="post" action="/api/register/confirm" onSubmit={onSubmit}>
					<label>
						<input type="hidden" name="email" ref={emailRef} required defaultValue={router.query?.email} />
					</label>
					<label>
						Confirmation Code:
						<input type="text" name="confirmCode" ref={codeRef} required />
					</label>
					<button type="submit" disabled={loading}>
						Confirm email
					</button>
				</form>
			</main>
		</>
	)
}
