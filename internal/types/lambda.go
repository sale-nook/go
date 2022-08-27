package types

type GraphQLQueryType string

const (
	GraphQLQueryTypeQuery        GraphQLQueryType = "Query"
	GraphQLQueryTypeMutation     GraphQLQueryType = "Mutation"
	GraphQLQueryTypeSubscription GraphQLQueryType = "Subscription"
)

type AppSyncLambdaIdentityEventIdentity struct {
	AccountID                   *string    `json:"accountId"`
	CognitoIdentityAuthProvider *string    `json:"cognitoIdentityAuthProvider"`
	CognitoIdentityID           *string    `json:"cognitoIdentityId"`
	CognitoIdentityPoolID       *string    `json:"cognitoIdentityPoolId"`
	SourceIP                    *[]*string `json:"sourceIp"`
	UserARN                     *string    `json:"userArn"`
	Username                    *string    `json:"username"`
	UserID                      *string    `json:"userId"`
}

type AppSyncLambdaIdentityEventInfo[Args any] struct {
	FieldName      *string           `json:"fieldName"`
	ParentTypeName *GraphQLQueryType `json:"parentTypeName"`
	Variables      *Args             `json:"variables"`
}

type AppSyncLambdaIdentityEvent[Args any] struct {
	Arguments *Args                                 `json:"arguments"`
	Identity  *AppSyncLambdaIdentityEventIdentity   `json:"identity"`
	Info      *AppSyncLambdaIdentityEventInfo[Args] `json:"info"`
	// I intentionally ignore the request object here to save memory and thus money.
}

type AppsyncLambdaResponseEvent[Type any] struct {
	Response *Type `json:"response"`
}

/*
THIS IS JUST A SAMPLE OF THE JSON OBJECT THAT IS PASSED TO THE LAMBDA.
	map[
		arguments:map[]
		identity:map[
			accountId:507348277062
			cognitoIdentityAuthProvider:"cognito-idp.eu-west-2.amazonaws.com/eu-west-2_OyqmXPUUb","cognito-idp.eu-west-2.amazonaws.com/eu-west-2_OyqmXPUUb:CognitoSignIn:6625e439-89a7-413b-b750-ef5e4cce38fa"
			cognitoIdentityAuthType:authenticated
			cognitoIdentityId:eu-west-2:86bfb2e2-035e-49b3-8336-7e4ed96d8a65
			cognitoIdentityPoolId:eu-west-2:0c342cb0-3528-40eb-b0d5-d2f91fd0bd9a
			sourceIp:[109.156.214.238]
			userArn:arn:aws:sts::507348277062:assumed-role/staging-appsync-stagingappsyncappsyncuser4D9308A6-1UP6WFQXVY46A/CognitoIdentityCredentials
			username:AROAXMICQK5DH5HLVHVSX:CognitoIdentityCredentials
		]
		info:map[
			fieldName:getProfile
			parentTypeName:Query
			selectionSetGraphQL:{}
			selectionSetList:[
				id
				jobs
				jobs/id
				jobs/status
				jobs/started
				jobs/ended
				jobs/__typename
				integrations
				integrations/id
				integrations/name
				integrations/keyValues
				integrations/keyValues/key
				integrations/keyValues/value
				integrations/keyValues/__typename
				integrations/__typename
				__typename
			]
			variables:map[]
		]
		prev:<nil>
		request:map[
			domainName:<nil>
			headers:map[
				accept:*\/*
				accept-encoding:gzip, deflate, br
				accept-language:en-GB,en-US;q=0.9,en;q=0.8
				authorization:AWS4-HMAC-SHA256 Credential=ASIAXMICQK5DJ7GVXOHB/20220807/eu-west-2/appsync/aws4_request, SignedHeaders=accept;content-type;host;x-amz-date;x-amz-security-token, Signature=39ce980678d3919e6d0bb7bf638e0967a1a3e984ec4b592404aa8e1b2b7adc6e
				cache-control:no-cache
				cloudfront-forwarded-proto:https
				cloudfront-is-desktop-viewer:true
				cloudfront-is-mobile-viewer:false
				cloudfront-is-smarttv-viewer:false
				cloudfront-is-tablet-viewer:false
				cloudfront-viewer-asn:2856
				cloudfront-viewer-country:GB
				content-length:383
				content-type:application/json;charset=UTF-8
				dnt:1
				host:it4wom7rnjdlvbvzu236urcmva.appsync-api.eu-west-2.amazonaws.com
				origin:http://localhost:3000
				pragma:no-cache
				referer:http://localhost:3000/
				sec-ch-ua:".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"
				sec-ch-ua-mobile:?0
				sec-ch-ua-platform:"macOS"
				sec-fetch-dest:empty
				sec-fetch-mode:cors
				sec-fetch-site:cross-site
				sec-gpc:1
				user-agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 via:2.0 76cca2ef798b9dc955bb151bf3bff218.cloudfront.net (CloudFront)
				x-amz-cf-id:sWka6vEqUvC9puZvYYQOzhQVtQmUQ8YhTBvAklLF5OikU0Q1DdC0kg== x-amz-date:20220807T145610Z
				x-amz-security-token:IQoJb3JpZ2luX2VjEJf//////////wEaCWV1LXdlc3QtMiJHMEUCICLd1EPRaFdJDP/tY/iexNs2ZXjOzSISAvhBW49wQpGhAiEAptWDbNgFRNvacNlROGI8UmoAvv8BuzWVD2JWekAQ1HgqzQQI8P//////////ARAAGgw1MDczNDgyNzcwNjIiDIuYKUzW5BN9nXKlYCqhBMpSLKGIZcWJJREVAK1oWzMzzw2YUgFTWop7nKZ68+B/XKFxYun2nt44v63Z/tkkzvW/1OwJT31HiGfno5KN7MGY/XZbxmx0bsSVUt9Gd41P1T5PfuEGGMqUaQXwiwttqyAxRsXbVpC5iy2bOFbVhAZGSbN5Q0+7af/vCtX85ZlarFYTVIgyV7XB2Tvn633+0ciK88od+WZsWfHTE/KY7ZQU/Olcieg7BumMqLm7y7wiUuqr5dl60Yb009syUY9lMQRLOz9LfrbG4Oa0M+e3quXpLguAsGnb/nyq3tnHzKTWjjyZ/J5l1iaOuGzVzlXWVkG7Uj+iTSOu4hYMFlW8xFR+Pi4Y4fgRf/52UqaRuS6So8jU/TPmlLHKOslHW4QrYWKnZU5Hu7wqBMUaAYZmNjKMamVlB8baEfo5z9N1Q8KBeW/v5TgJwnHKspntCAhZhMj75U/APl/X9HPcckCnNCqNJTaQpbg6NLwjT9xJVS47vp9/UZkP1MRH0069xkd1oO8fVm877Gkl5NfNYg93fPCsSsjEWLdEu1ZFv4pixL+I10arEwHa9UX/bs/n0OMUe45s4i5jdQ0x9riqJNL0AAkDGmSBAAyJBrUGD9Qd/OsoF6UrF6DlRmhoCEzkuJGCCR7Kx3KLTMsUEgkPaMEiYM5KwcVMoSgR949LwlmkVfXC9G1v+o0KElLXWVWRutCOn+VWz9zjt0ZxUh6Eto0jdvnHMIqlv5cGOoUCMPtaoaxTcVPIpGfisI17ers+g8VJRcJ9DHkhzbFEwcpOQGNCulknDLTqgPipJHYyFajSj2kuGjxwPKviznwLr0eBvwXsW9SAcg8EGW3gqHnFGFLfE4ABl2OzjzLo83g+LUTvq25K4Mfk4yTm8Dd+eeCSv5jRqUv4N4stiP0vdmt2ZFh4s+Z7iAqiXOe2mAThdHkw94JKu6Z6vMF/IrwiuB7+9MXgMWOejwaW6gF61781vUWaQe5rF64qRpRB0Uq04B73qk996xxgtz9/ecgBl06BfdiXFucVd8sgwvHEwLne1xDbCTzYRlbvgUO5EDqMGtsnL2Qdq6NP1ID6DCaTfPdyKNL0
				x-amz-user-agent:aws-amplify/2.0.8
				x-amzn-requestid:63ea06bc-db65-4aaa-98ae-978c928b23de
				x-amzn-trace-id:Root=1-62efd28b-0153aa8538ee4d13136b21a9
				x-forwarded-for:109.156.214.238, 15.158.16.49
				x-forwarded-port:443
				x-forwarded-proto:https
			]
		]
		source:<nil>
		stash:map[]
	]*/
