package internal

type App struct {
	Name string `json:"name"`
}

func GetApp() App {
	return App{
		Name: "cdk-appsync-go",
	}
}
