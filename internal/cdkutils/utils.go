package cdkutils

import (
	"fmt"
	"os"
	"strings"

	"github.com/davemackintosh/go/internal"
)

func trimAllPrefixes(name, prefix string) string {
	out := name
	for {
		if !strings.HasPrefix(out, prefix) {
			break
		}
		out = strings.TrimPrefix(out, prefix)
	}

	return out
}

func getLoweredEnvironment() string {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "staging"
	}

	return strings.ToLower(strings.TrimSpace(environment))
}

func NameWithEnvironment(name string) string {
	environmentLowered := getLoweredEnvironment()
	appMeta := internal.GetApp()

	// Trim off any existing environment prefixes
	name = trimAllPrefixes(name, appMeta.Name+"-")
	name = trimAllPrefixes(name, environmentLowered+"-")

	return fmt.Sprintf("%s-%s-%s", appMeta.Name, environmentLowered, name)
}

func NameWithStackAndEnvironment(name, stackName string) string {
	environmentLowered := getLoweredEnvironment()
	stackName = trimAllPrefixes(stackName, stackName+"-")
	appMeta := internal.GetApp()

	// Trim off any existing environment prefixes
	name = trimAllPrefixes(name, appMeta.Name+"-")
	name = trimAllPrefixes(name, environmentLowered+"-")

	return NameWithEnvironment(fmt.Sprintf("%s-%s-%s", appMeta.Name, stackName, name))
}
