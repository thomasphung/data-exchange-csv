# Testing internal/server/http.go using behavior-driven testing methodology
Feature: Use /health endpoint
    In order to verify that I can access the API
    I need to be able to connect to /health endpoint

    Scenario Outline: Valid request
        When I make a <http_method> request
        Then it should return a 200 HTTP status

        Examples:
            | http_method |
            | GET |
    
    Scenario Outline: Invalid request
        When I make a <http_method> request
        Then it should return a 405 HTTP status

        Examples:
            | http_method |
            | CONNECT |
            | DELETE |
            | HEAD |
            | OPTIONS |
            | POST |
            | PUT |
            | TRACE |
