import { describe, test, expect, vi } from "vitest"
import { getApiConfig } from "./api"

vi.mock("../../../config/api.json", () => ({
	default: {
		"go-staging-appsync": {
			UserPoolId: "go-staging-userpool",
			UserPoolClientId: "go-staging-userpool-client",
			IdentityPoolId: "go-staging-identitypool",
			AppSyncURL: "go-staging-appsync.appsync-api.eu-west-1.amazonaws.com",
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
