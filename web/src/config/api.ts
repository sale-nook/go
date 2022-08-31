import apiConfig from "../../../config/api.json"

export interface ApiConfig {
	IdentityPoolID: string
	UserPoolID: string
	UserPoolClientID: string
	AppsyncURL: string
	Region: string
}

export enum configKey {
	Staging = "cdk-appsync-go-staging-appsync",
}

function mustBeSet<R extends Object>(debugName: string, value?: R): R {
	if (!value || value.toString() === "") throw new Error(`${debugName} is not set`)

	return value
}

export function getApiConfig(): ApiConfig {
	const env = mustBeSet("ENVIRONMENT", process.env.ENVIRONMENT)

	const confKey = `cdk-appsync-go-${env}-appsync` as configKey
	// If you're seeing a TypeError here, you haven't yet deployed the application to
	// the environment you're testing against. You can do that by running `deploy`.
	// The CDK outputs the config json to a file in `<project-root>/config/api.json`.
	const conf = mustBeSet<typeof apiConfig["cdk-appsync-go-staging-appsync"]>(confKey, apiConfig[confKey])

	return {
		UserPoolID: mustBeSet("UserPoolID", conf.UserPoolId),
		UserPoolClientID: mustBeSet("UserPoolClientID", conf.UserPoolClientId),
		IdentityPoolID: mustBeSet("IdentityPoolID", conf.IdentityPoolId),
		AppsyncURL: mustBeSet("AppsyncURL", conf.AppSyncURL),
		Region: mustBeSet("AWSRegion", conf.Region),
	}
}
