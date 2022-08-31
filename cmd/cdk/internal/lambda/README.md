# Lambda CDK Helper

## Types

### FunctionProps

```go
type FunctionProps struct {
	[RuntimeEnvironmentals](#RuntimeEnvironmentals) *map[string]*string
	[Entry](#Entry)                *string
	[API](#API)                   *awsappsync.CfnGraphQLApi
	[Vpc](#VPC)                   *awsec2.Vpc
	[URLProps](#URLProps)              *awslambda.FunctionUrlOptions
}
```

###### `RuntimeEnvironmentals`

A map of environmental variables this Lambda should execute with. As with all other parts of this library, `ENVIRONMENT` is always set to the current environment and overriden if set here.

###### `Entry`

This is the name of the compiled binary file from `go build` that the lambda will run.

###### `API`

This is the `CfnGraphQLApi` that this lambda will be a datasource for. This implies a dependency on the stack that the api is defined in.

###### `Vpc`

_currently_ unused. Otherwise a valid AWS VPC instance this lambda belongs to.

###### `URLProps`

Props to pass to the `lambda.AddFunctionUrl` function. Not supplying this property negates the creation of a function url.

## Functions

### `NewLambdaFunction(name string, stack awscdk.Stack, props *FunctionProps) awslambda.Function`

Create a new lambda function with the provided `props` of type [`FunctionProps`](#FunctionProps)
