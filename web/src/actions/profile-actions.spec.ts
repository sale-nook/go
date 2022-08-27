import { describe, expect, test } from "vitest"
import { SetProfileAction, initProfileAction } from "./profile-actions"
import { ProfileActionTypes } from "src/reducers/profile"

describe("Profile Actions", () => {
	test("should create a profile", () => {
		expect(initProfileAction()).toEqual({
			type: ProfileActionTypes.INIT,
		})
	})

	test("should set profile", () => {
		const profile = {
			id: 1,
			jobs: [],
			integrations: [],
		}
		expect(SetProfileAction(profile)).toEqual({
			type: ProfileActionTypes.SET_PROFILE,
			payload: profile,
		})
	})
})
