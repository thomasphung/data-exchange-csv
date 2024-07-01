# Testing internal/server/http.go using behavior-driven testing methodology
Feature: Use /health endpoint
    In order to verify that I can access the API
    I need to be able to connect to /health endpoint

    Scenario: GET request
        Given I make a GET request to /health
        Then it should return a 200 HTTP status
