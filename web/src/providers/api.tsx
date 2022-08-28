import { gql } from "@apollo/client"
import * as Apollo from "@apollo/client"
export type Maybe<T> = T | null
export type InputMaybe<T> = Maybe<T>
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] }
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> }
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> }
const defaultOptions = {} as const
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
	ID: string
	String: string
	Boolean: boolean
	Int: number
	Float: number
	AWSDateTime: any
}

/**
 * When registering a user, a verification email is sent to the user's email address
 * with a code that must be entered to complete the registration.
 */
export type ConfirmEmailInput = {
	code: Scalars["String"]
	email: Scalars["String"]
}

export type ConfirmEmailOutput = {
	__typename?: "ConfirmEmailOutput"
	email?: Maybe<Scalars["String"]>
}

export type LoginInput = {
	email: Scalars["String"]
	password: Scalars["String"]
}

export type Mutation = {
	__typename?: "Mutation"
	confirmEmail?: Maybe<ConfirmEmailOutput>
	login?: Maybe<User>
	register?: Maybe<RegisterOutput>
}

export type MutationConfirmEmailArgs = {
	input: ConfirmEmailInput
}

export type MutationLoginArgs = {
	input: LoginInput
}

export type MutationRegisterArgs = {
	input: RegisterInput
}

export type PlaidLinkOutput = {
	__typename?: "PlaidLinkOutput"
	link?: Maybe<Scalars["String"]>
}

export type Profile = {
	__typename?: "Profile"
	createdAt: Scalars["AWSDateTime"]
	createdBy: Scalars["String"]
	email: Scalars["String"]
	id: Scalars["String"]
	name: Scalars["String"]
}

export type Query = {
	__typename?: "Query"
	getProfile?: Maybe<User>
	initiatePlaidLink?: Maybe<PlaidLinkOutput>
}

export type RegisterInput = {
	email: Scalars["String"]
	password: Scalars["String"]
	repeatPassword: Scalars["String"]
}

export type RegisterOutput = {
	__typename?: "RegisterOutput"
	userID?: Maybe<Scalars["String"]>
}

export type User = {
	__typename?: "User"
	createdAt?: Maybe<Scalars["AWSDateTime"]>
	id: Scalars["String"]
	profiles?: Maybe<Array<Maybe<Profile>>>
}

export type ConfirmEmailMutationVariables = Exact<{
	input: ConfirmEmailInput
}>

export type ConfirmEmailMutation = {
	__typename?: "Mutation"
	confirmEmail?: { __typename?: "ConfirmEmailOutput"; email?: string | null } | null
}

export type FullUserFragment = {
	__typename?: "User"
	id: string
	createdAt?: any | null
	profiles?: Array<{
		__typename?: "Profile"
		id: string
		createdAt: any
		createdBy: string
		name: string
		email: string
	} | null> | null
}

export type GetProfileQueryVariables = Exact<{ [key: string]: never }>

export type GetProfileQuery = {
	__typename?: "Query"
	getProfile?: {
		__typename?: "User"
		id: string
		createdAt?: any | null
		profiles?: Array<{
			__typename?: "Profile"
			id: string
			createdAt: any
			createdBy: string
			name: string
			email: string
		} | null> | null
	} | null
}

export type LoginMutationVariables = Exact<{
	input: LoginInput
}>

export type LoginMutation = {
	__typename?: "Mutation"
	login?: {
		__typename?: "User"
		id: string
		createdAt?: any | null
		profiles?: Array<{
			__typename?: "Profile"
			id: string
			createdAt: any
			createdBy: string
			name: string
			email: string
		} | null> | null
	} | null
}

export type RegisterNewUserMutationVariables = Exact<{
	input: RegisterInput
}>

export type RegisterNewUserMutation = {
	__typename?: "Mutation"
	register?: { __typename?: "RegisterOutput"; userID?: string | null } | null
}

export const FullUserFragmentDoc = gql`
	fragment FullUser on User {
		id
		profiles {
			id
			createdAt
			createdBy
			name
			email
		}
		createdAt
	}
`
export const ConfirmEmailDocument = gql`
	mutation ConfirmEmail($input: ConfirmEmailInput!) {
		confirmEmail(input: $input) {
			email
		}
	}
`
export type ConfirmEmailMutationFn = Apollo.MutationFunction<ConfirmEmailMutation, ConfirmEmailMutationVariables>

/**
 * __useConfirmEmailMutation__
 *
 * To run a mutation, you first call `useConfirmEmailMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useConfirmEmailMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [confirmEmailMutation, { data, loading, error }] = useConfirmEmailMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useConfirmEmailMutation(
	baseOptions?: Apollo.MutationHookOptions<ConfirmEmailMutation, ConfirmEmailMutationVariables>,
) {
	const options = { ...defaultOptions, ...baseOptions }
	return Apollo.useMutation<ConfirmEmailMutation, ConfirmEmailMutationVariables>(ConfirmEmailDocument, options)
}
export type ConfirmEmailMutationHookResult = ReturnType<typeof useConfirmEmailMutation>
export type ConfirmEmailMutationResult = Apollo.MutationResult<ConfirmEmailMutation>
export type ConfirmEmailMutationOptions = Apollo.BaseMutationOptions<
	ConfirmEmailMutation,
	ConfirmEmailMutationVariables
>
export const GetProfileDocument = gql`
	query getProfile {
		getProfile {
			...FullUser
		}
	}
	${FullUserFragmentDoc}
`

/**
 * __useGetProfileQuery__
 *
 * To run a query within a React component, call `useGetProfileQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetProfileQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetProfileQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetProfileQuery(baseOptions?: Apollo.QueryHookOptions<GetProfileQuery, GetProfileQueryVariables>) {
	const options = { ...defaultOptions, ...baseOptions }
	return Apollo.useQuery<GetProfileQuery, GetProfileQueryVariables>(GetProfileDocument, options)
}
export function useGetProfileLazyQuery(
	baseOptions?: Apollo.LazyQueryHookOptions<GetProfileQuery, GetProfileQueryVariables>,
) {
	const options = { ...defaultOptions, ...baseOptions }
	return Apollo.useLazyQuery<GetProfileQuery, GetProfileQueryVariables>(GetProfileDocument, options)
}
export type GetProfileQueryHookResult = ReturnType<typeof useGetProfileQuery>
export type GetProfileLazyQueryHookResult = ReturnType<typeof useGetProfileLazyQuery>
export type GetProfileQueryResult = Apollo.QueryResult<GetProfileQuery, GetProfileQueryVariables>
export const LoginDocument = gql`
	mutation Login($input: LoginInput!) {
		login(input: $input) {
			...FullUser
		}
	}
	${FullUserFragmentDoc}
`
export type LoginMutationFn = Apollo.MutationFunction<LoginMutation, LoginMutationVariables>

/**
 * __useLoginMutation__
 *
 * To run a mutation, you first call `useLoginMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useLoginMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [loginMutation, { data, loading, error }] = useLoginMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useLoginMutation(baseOptions?: Apollo.MutationHookOptions<LoginMutation, LoginMutationVariables>) {
	const options = { ...defaultOptions, ...baseOptions }
	return Apollo.useMutation<LoginMutation, LoginMutationVariables>(LoginDocument, options)
}
export type LoginMutationHookResult = ReturnType<typeof useLoginMutation>
export type LoginMutationResult = Apollo.MutationResult<LoginMutation>
export type LoginMutationOptions = Apollo.BaseMutationOptions<LoginMutation, LoginMutationVariables>
export const RegisterNewUserDocument = gql`
	mutation RegisterNewUser($input: RegisterInput!) {
		register(input: $input) {
			userID
		}
	}
`
export type RegisterNewUserMutationFn = Apollo.MutationFunction<
	RegisterNewUserMutation,
	RegisterNewUserMutationVariables
>

/**
 * __useRegisterNewUserMutation__
 *
 * To run a mutation, you first call `useRegisterNewUserMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useRegisterNewUserMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [registerNewUserMutation, { data, loading, error }] = useRegisterNewUserMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useRegisterNewUserMutation(
	baseOptions?: Apollo.MutationHookOptions<RegisterNewUserMutation, RegisterNewUserMutationVariables>,
) {
	const options = { ...defaultOptions, ...baseOptions }
	return Apollo.useMutation<RegisterNewUserMutation, RegisterNewUserMutationVariables>(
		RegisterNewUserDocument,
		options,
	)
}
export type RegisterNewUserMutationHookResult = ReturnType<typeof useRegisterNewUserMutation>
export type RegisterNewUserMutationResult = Apollo.MutationResult<RegisterNewUserMutation>
export type RegisterNewUserMutationOptions = Apollo.BaseMutationOptions<
	RegisterNewUserMutation,
	RegisterNewUserMutationVariables
>
