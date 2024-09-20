package main

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

// tagging in go
type MyEvent struct {
	Username string `json:"username"`
}

// handler
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	// returns a value for us to use and nil for error
	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

// backend logic for application
// packaged, zipped and deployed to our lambda function
func main() {
	lambdaApp := app.NewApp()
	lambda.Start(lambdaApp.ApiHandler.RegisterUserHandler)
}
