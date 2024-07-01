package main

import (
	"testing"

	"github.com/cucumber/godog"
)

// godogsCtxKey is the key used to store the available godogs in the context.Context.
type godogsCtxKey struct{}

func iMakeGETRequest(arg1 int) error {
	return godog.ErrPending
}

func itShouldReturnStatus(arg1 int) error {
	return godog.ErrPending
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I make ([A-Z]+) request$`, iMakeGETRequest) // TODO: Better define HTTP verbs
	ctx.Step(`^it should return a (\d+) HTTP status$`, itShouldReturnStatus)
}
