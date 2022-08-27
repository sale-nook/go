import { describe, expect, test } from "vitest"
import { AppRoutePaths, homeRoute, loginRoute, registerConfirmRoute, registerRoute } from "./routes"

describe("Routes", () => {
	test("Routes should be typesafe.", () => {
		expect(homeRoute()).toEqual(AppRoutePaths.Home)
		expect(loginRoute()).toEqual(AppRoutePaths.Login)
		expect(registerRoute()).toEqual(AppRoutePaths.Register)
		expect(
			registerConfirmRoute({
				query: {
					email: "aws-appsync-go@aws-appsync-go.io",
				},
			}),
		).toEqual({
			pathname: AppRoutePaths.RegisterConfirm,
			query: {
				email: "aws-appsync-go%40aws-appsync-go.io",
			},
		})
		expect(registerConfirmRoute()).toEqual(AppRoutePaths.RegisterConfirm)
	})
})
