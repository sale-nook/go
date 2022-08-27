import {
	InitActionPayload,
	ProfileActionTypes,
	SetProfileActionActionPayload,
	SetProfileErrorActionPayload,
} from "src/reducers/profile"
import { User } from "src/providers/api"

/**
 * Initialize the profile state
 */
export function initProfileAction(): InitActionPayload {
	return {
		type: ProfileActionTypes.INIT,
	}
}

/**
 * Update or remove the profile in the profile context,
 * also removes the error message and cookies.
 */
export function SetProfileAction(profile: User | null): SetProfileActionActionPayload {
	return {
		type: ProfileActionTypes.SET_PROFILE,
		payload: profile,
	}
}

/**
 * Set or remove the error message for the profile context.
 */
export function SetProfileErrorAction(error: string | null): SetProfileErrorActionPayload {
	return {
		type: ProfileActionTypes.SET_PROFILE_ERROR,
		payload: error,
	}
}
