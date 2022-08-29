import { describe, test, expect, vi } from "vitest"
import { getApiConfig } from "./api"

vi.mock("../../../config/api.json", () => ({
	default: {
		"aws-appsync-go-staging-appsync": {
			UserPoolId: "aws-appsync-go-staging-userpool",
			UserPoolClientId: "aws-appsync-go-staging-userpool-client",
			IdentityPoolId: "aws-appsync-go-staging-identitypool",
			AppSyncURL: "aws-appsync-go-staging-appsync.appsync-api.eu-west-1.amazonaws.com",
		},
	},
}))

describe("Get api config object", () => {
	test("throws expected errors", () => {
		expect(() => {
			process.env.ENVIRONMENT = ""
			getApiConfig()
		}).toThrowError("ENVIRONMENT is not set")

		expect(() => {
			process.env.ENVIRONMENT = "nope"
			process.env.USER_POOL_CLIENT_SECRET = "staging"
			getApiConfig()
		}).toThrowError("nope-appsync is not set")
	})
})
