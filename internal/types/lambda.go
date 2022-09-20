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

