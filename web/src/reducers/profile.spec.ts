import { ProfileActionTypes, profileReducer as profileReducer } from "./profile"
import { describe, expect, test } from "vitest"

describe("Profile reducer tests", () => {
	test("returns default state", () => {
		const state = profileReducer(undefined, { type: ProfileActionTypes.INIT })
		expect(state).toEqual({
			profile: null,
			error: null,
		})
	})

	test("returns state with profile", () => {
		const state = profileReducer(undefined, {
			type: ProfileActionTypes.SET_PROFILE,
			payload: {
				id: "123",
				integrations: [],
				jobs: [],
			},
		})
		expect(state).toEqual({
			profile: {
				id: "123",
				integrations: [],
				jobs: [],
			},
			error: null,
		})
	})

	test("returns state with error", () => {
		const state = profileReducer(undefined, {
			type: ProfileActionTypes.SET_PROFILE_ERROR,
			payload: "error",
		})
		expect(state).toEqual({
			profile: null,
			error: "error",
		})
	})
})
