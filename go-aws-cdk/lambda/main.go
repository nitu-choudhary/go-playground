package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// tagging in go
type MyEvent struct {
	Username string `json:"username"`
}

// simple func that takes in payload and performs some action
// error is native type in Go
// errors are values in go - not like try catch or exceptions like in other languages
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
	lambda.Start(HandleRequest)
}
