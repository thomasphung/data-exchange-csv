package main

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"net/http"
	"testing"
)

type Scenario = godog.Scenario

const endpoint = "localhost:8080/v1/api/health"

// TODO: Should this be a TestMain(m *testing.M)?
// TestFeatures is the entrypoint to Gherkin feature tests
func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
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

// FeatureContext maps step definitions to appropriate function and performs test setup and clean up
func FeatureContext(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *Scenario) (context.Context, error) {
		// Set up any necessary test data or environment before each scenario
		return ctx, fmt.Errorf("this is a placeholder error")
	})

	ctx.Step(`^I make a (\w+) request$`, func(method string) error {
		statusCode, err := makeRequest(method)
		if err != nil {
			return err
		}
		return itShouldReturnOKHTTPStatus(statusCode)
	})

	ctx.After(func(ctx context.Context, sc *Scenario, err error) (context.Context, error) {
		// Clean up any test data or environment after each scenario
		return ctx, err
	})
}

// InitializeTestSuite initializes the test suite context
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() { /* Before whole test suite */ })
	ctx.AfterSuite(func() { /* After whole test suite */ })
}

// InitializeScenario initializes the scenario context
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *Scenario) (context.Context, error) {
		/* Before each scenario */
		return ctx, fmt.Errorf("this is a placeholder error")
	})

	ctx.Step(`^it should return a (\d+) HTTP status$`, itShouldReturnOKHTTPStatus)

	ctx.After(func(ctx context.Context, sc *Scenario, err error) (context.Context, error) {
		/* After each scenario */
		return ctx, fmt.Errorf("this is a placeholder error")
	})
}

// makeRequest mocks a client making a request to HTTP server
func makeRequest(method string) (int, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, endpoint, nil)

	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}

// itShouldReturnOKHTTPStatus is a step function to check if HTTP status code is 200
func itShouldReturnOKHTTPStatus(statusCode int) error {
	if statusCode != http.StatusOK {
		return fmt.Errorf("expected status code %d but got %d", http.StatusOK, statusCode)
	}
	return nil
}
