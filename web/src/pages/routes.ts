import { useRouter } from "next/router"
import { ParsedUrlQuery } from "querystring"
import { UrlObject } from "url"

export enum AppRoutePaths {
	Home = "/",
	Login = "/login",
	ForgotPassword = "/login/forgot-password",
	Logout = "/logout",
	Register = "/register",
	ResetPassword = "/register/reset-password",
	RegisterConfirm = "/register/confirm",
}

export interface RegisterConfirmParams {
	email: string
}

export interface OptionalEmailParams {
	email?: string
}

interface RouteOptions<Params, Query> {
	params?: Params
	query?: Query
}

interface Route<Params extends object, Query extends object> {
	path: AppRoutePaths
	options?: RouteOptions<Params, Query>
}

interface HomeRoute extends Route<never, never> {
	path: AppRoutePaths.Home
}

interface LoginRoute extends Route<never, OptionalEmailParams> {
	path: AppRoutePaths.Login
}

interface LogoutRoute extends Route<never, never> {
	path: AppRoutePaths.Logout
}

interface RegisterRoute extends Route<never, never> {
	path: AppRoutePaths.Register
}

interface RegisterConfirmRoute extends Route<never, RegisterConfirmParams> {
	path: AppRoutePaths.RegisterConfirm
}

interface ForgotPasswordRoute extends Route<never, OptionalEmailParams> {
	path: AppRoutePaths.ForgotPassword
}

interface ResetPasswordRoute extends Route<never, RegisterConfirmParams> {
	path: AppRoutePaths.ResetPassword
}

export type Routes =
	| HomeRoute
	| LoginRoute
	| LogoutRoute
	| RegisterRoute
	| RegisterConfirmRoute
	| ForgotPasswordRoute
	| ResetPasswordRoute

function routeParamIsStringable(value: any): value is string {
	return typeof value === "string" || typeof value === "number"
}

/**
 * Get the /type safe/ query parameters of this route.
 */
export function useQueryParams<T extends object>(): Record<keyof T, string> {
	const router = useRouter()
	const query = router.query as ParsedUrlQuery
	const params = {} as Record<keyof T, string>

	for (const [key, value] of Object.entries(query)) {
		params[key as keyof T] = decodeURIComponent(value as string)
	}

	return params
}

/**
 * Compiles a route into a URL that can be used to push to.
 * Supports routes with params and query strings.
 *
 * @param route to compile
 * @returns string
 */
function compileRoute(route: Routes): UrlObject | string {
	let routePath: string = route.path

	if (route.options?.params) {
		for (const [key, value] of Object.entries(route.options.params)) {
			if (routeParamIsStringable(value)) {
				routePath = routePath.replace(`:${key}`, encodeURIComponent(value))
			}
		}
	}

	if (route.options?.query) {
		const query: ParsedUrlQuery = {}

		for (const [key, value] of Object.entries(route.options.query)) {
			if (routeParamIsStringable(value)) {
				query[key] = encodeURIComponent(value)
			} else {
				console.warn(`Query param ${key} is not stringable, this could be dangerous.`)
				query[key] = value
			}
		}

		return {
			pathname: routePath,
			query,
		}
	}

	return routePath
}

export function homeRoute() {
	return compileRoute({
		path: AppRoutePaths.Home,
	})
}

export function loginRoute(email?: string) {
	return compileRoute({
		path: AppRoutePaths.Login,
		options: {
			query: {
				email,
			},
		},
	})
}

export function forgotPasswordRoute(email?: string) {
	return compileRoute({
		path: AppRoutePaths.ForgotPassword,
		options: {
			query: {
				email,
			},
		},
	})
}

export function logoutRoute() {
	return compileRoute({
		path: AppRoutePaths.Logout,
	})
}

export function registerRoute() {
	return compileRoute({
		path: AppRoutePaths.Register,
	})
}

export function registerConfirmRoute(email: string) {
	return compileRoute({
		path: AppRoutePaths.RegisterConfirm,
		options: {
			query: {
				email,
			},
		},
	})
}

export function resetPasswordRoute(email: string) {
	return compileRoute({
		path: AppRoutePaths.ResetPassword,
		options: {
			query: {
				email,
			},
		},
	})
}
