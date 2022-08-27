import apiConfig from "../../../config/api.json"

export interface ApiConfig {
	UserPoolClientID: string
	UserPoolClientSecret: string
	IdentityPoolID: string
	UserPoolID: string
	AppsyncURL: string
	Region: string
}

export enum configKey {
	Staging = "aws-appsync-go-staging-appsync",
}

function mustBeSet<R extends Object>(name: string, value?: R): R {
	if (!value || value.toString() === "") throw new Error(`${name} is not set`)

	return value
}

export function getApiConfig(): ApiConfig {
	const env = mustBeSet("ENVIRONMENT", process.env.ENVIRONMENT)

	const confKey = `aws-appsync-go-${env}-appsync` as configKey
	const conf = mustBeSet<typeof apiConfig["aws-appsync-go-staging-appsync"]>(confKey, apiConfig[confKey])

	return {
		UserPoolClientID: mustBeSet("UserPoolClientID", conf.UserPoolClientId),
		UserPoolID: mustBeSet("UserPoolID", conf.UserPoolId),
		IdentityPoolID: mustBeSet("IdentityPoolID", conf.IdentityPoolId),
		AppsyncURL: mustBeSet("AppsyncURL", conf.AppSyncURL),
		Region: mustBeSet("AWS_REGION", process.env.AWS_REGION),
		UserPoolClientSecret: mustBeSet("USER_POOL_CLIENT_SECRET", process.env.USER_POOL_CLIENT_SECRET),
	}
}
