package types

type CDKOutputs struct {
	UserPoolClientID *string `json:"goUserPoolClientId"` //nolint: tagliatelle
	UserPoolID       *string `json:"goUserPoolId"`       //nolint: tagliatelle
}

type CDKOutputsByEnv struct {
	Staging    CDKOutputs `json:"go-staging-appsync"` //nolint: tagliatelle
	Production CDKOutputs `json:"go-prod-appsync"`    //nolint: tagliatelle
}
