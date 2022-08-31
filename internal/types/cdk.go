package types

type CDKOutputs struct {
	UserPoolClientID *string `json:"cdk-appsync-goUserPoolClientId"` //nolint: tagliatelle
	UserPoolID       *string `json:"cdk-appsync-goUserPoolId"`       //nolint: tagliatelle
}

type CDKOutputsByEnv struct {
	Staging    CDKOutputs `json:"cdk-appsync-go-staging-appsync"` //nolint: tagliatelle
	Production CDKOutputs `json:"cdk-appsync-go-prod-appsync"`    //nolint: tagliatelle
}
