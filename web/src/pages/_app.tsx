import { NextComponentType, NextPageContext } from "next"
import type { AppProps } from "next/app"
import { useRouter } from "next/router"
import { PropsWithChildren, useEffect } from "react"
import { ApolloWrapper } from "src/providers/apollo"
import { UserProvider, useUser } from "src/providers/user"
import "../styles/global.css"

export default function App({ Component, pageProps: { session, ...pageProps } }: AppProps) {
	return (
		<ApolloWrapper>
			<UserProvider>
				{(Component as NextComponentType<NextPageContext, any, {}> & { auth: boolean }).auth ? (
					<Auth>
						<Component {...pageProps} />
					</Auth>
				) : (
					<Component {...pageProps} />
				)}
			</UserProvider>
		</ApolloWrapper>
	)
}

function Auth({ children }: PropsWithChildren<{}>) {
	const { user, loading } = useUser()
	const router = useRouter()

	useEffect(() => {
		if (!user) {
			router.push("/login")
		}
	}, [user])

	if (loading) {
		return <div>Loading...</div>
	}

	return <>{children}</>
}
