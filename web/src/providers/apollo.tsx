import { createAuthLink, AUTH_TYPE } from "aws-appsync-auth-link"
import { createSubscriptionHandshakeLink } from "aws-appsync-subscription-link"
import { ApolloProvider, ApolloClient, InMemoryCache, HttpLink, ApolloLink } from "@apollo/client"
import { getApiConfig } from "src/config/api"
import { PropsWithChildren, useEffect } from "react"
import { Auth } from "@aws-amplify/auth"

export const ApolloWrapper = ({ children }: PropsWithChildren<unknown>) => {
	const config = getApiConfig()
	const url = config.AppsyncURL
	const region = config.Region

	// You win this round Amplify.
	// I will remove you in the future though you are a bad person.
	useEffect(() => {
		Auth.configure({
			Auth: {
				region,
				userPoolId: config.UserPoolID,
				userPoolWebClientId: config.UserPoolClientID,
			},
			aws_appsync_graphqlEndpoint: url,
			aws_appsync_region: region,
			aws_appsync_authenticationType: AUTH_TYPE.AWS_IAM,
			aws_cognito_identity_pool_id: config.IdentityPoolID,
		})
	}, [])

	const auth = {
		url,
		region,
		auth: {
			url,
			type: AUTH_TYPE.AWS_IAM,
			credentials: async () => await Auth.currentCredentials(),
		},
	}

	const httpLink = new HttpLink({ uri: url })

	const link = ApolloLink.from([
		// @ts-ignore
		createAuthLink(auth),
		// @ts-ignore
		createSubscriptionHandshakeLink({ url, region, auth }, httpLink),
	])

	const client = new ApolloClient({
		link,
		cache: new InMemoryCache(),
	})

	return <ApolloProvider client={client}>{children}</ApolloProvider>
}
