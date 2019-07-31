Feature: Using the Lambda Handler
    Scenario: Getting a 200
        Given I send a GET request
        Then I receive a response with: "hello world"
