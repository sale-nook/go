import { describe, test, expect } from "vitest"
import { getApiConfig } from "./api"

describe("Get api config object", () => {
	test("throws expected errors", () => {
		expect(() => {
			process.env.ENVIRONMENT = ""
			getApiConfig()
		}).toThrowError("ENVIRONMENT is not set")

		expect(() => {
			process.env.ENVIRONMENT = "staging"
			process.env.USER_POOL_CLIENT_SECRET = ""
			getApiConfig()
		}).toThrowError("USER_POOL_CLIENT_SECRET is not set")

		expect(() => {
			process.env.ENVIRONMENT = "staging"
			process.env.USER_POOL_CLIENT_SECRET = "staging"
			process.env.AWS_REGION = ""
			getApiConfig()
		}).toThrowError("AWS_REGION is not set")

		expect(() => {
			process.env.ENVIRONMENT = "nope"
			process.env.USER_POOL_CLIENT_SECRET = "staging"
			process.env.AWS_REGION = "eu-west-2"
			getApiConfig()
		}).toThrowError("nope-appsync is not set")
	})
})
