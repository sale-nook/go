import { User } from "../providers/api"

export enum ProfileActionTypes {
	INIT = "@@profile/INIT",
	SET_PROFILE = "@@profile/SET_PROFILE",
	SET_PROFILE_ERROR = "@@profile/SET_PROFILE_ERROR",
}

interface ProfileActionBase {
	type: ProfileActionTypes
}

export interface SetProfileActionActionPayload extends ProfileActionBase {
	type: ProfileActionTypes.SET_PROFILE
	payload: User | null
}

export interface SetProfileErrorActionPayload extends ProfileActionBase {
	type: ProfileActionTypes.SET_PROFILE_ERROR
	payload: string | null
}

export interface InitActionPayload extends ProfileActionBase {
	type: ProfileActionTypes.INIT
}

type ProfileAction = InitActionPayload | SetProfileActionActionPayload | SetProfileErrorActionPayload

export interface ProfileState {
	profile: User | null
	error: string | null
}

const initialState: ProfileState = {
	profile: null,
	error: null,
}

export function profileReducer(state = initialState, action: ProfileAction) {
	switch (action.type) {
		case ProfileActionTypes.INIT:
			return initialState
		case ProfileActionTypes.SET_PROFILE:
			return {
				...state,
				profile: action.payload,
			}
		case ProfileActionTypes.SET_PROFILE_ERROR:
			return {
				...state,
				error: action.payload,
			}
		default:
			return state
	}
}
