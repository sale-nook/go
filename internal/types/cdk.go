package types

type CDKOutputs struct {
	UserPoolClientID *string `json:"aws-appsync-goUserPoolClientId"`
	UserPoolID       *string `json:"aws-appsync-goUserPoolId"`
}

type CDKOutputsByEnv struct {
	Staging    CDKOutputs `json:"aws-appsync-go-staging-appsync"` // nolint: tagliatelle
	Production CDKOutputs `json:"aws-appsync-go-prod-appsync"`    // nolint: tagliatelle
}
