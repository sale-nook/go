import { describe, expect, test } from "vitest"
import { AppRoutePaths, homeRoute, loginRoute, registerConfirmRoute, registerRoute } from "./routes"

describe("Routes", () => {
	test("Routes should be typesafe.", () => {
		const email = "aws-appsync-go@aws-appsync-go.io"
		expect(homeRoute()).toEqual(AppRoutePaths.Home)
		expect(loginRoute()).toEqual({
			pathname: AppRoutePaths.Login,
			query: {},
		})
		expect(registerRoute()).toEqual(AppRoutePaths.Register)
		expect(registerConfirmRoute(email)).toEqual({
			pathname: AppRoutePaths.RegisterConfirm,
			query: {
				email: encodeURIComponent(email),
			},
		})
		expect(registerConfirmRoute(email)).toEqual({
			pathname: AppRoutePaths.RegisterConfirm,
			query: {
				email: encodeURIComponent(email),
			},
		})
	})
})
