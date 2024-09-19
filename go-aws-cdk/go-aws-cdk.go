package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GoAwsCdkStackProps struct {
	awscdk.StackProps
}

func NewGoAwsCdkStack(scope constructs.Construct, id string, props *GoAwsCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// lambda resource
	// the passed stack is going to have the following lambda function with given name and function properties
	// every resource will have its own definition of properties
	// runtime - what is the runtime for code? go, java, nodejs etc
	// handler is where out code is packaged and bundled
	// code - where and how the code is being executed?
	awslambda.NewFunction(stack, jsii.String("goPlaygroundLambdaFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("lambda/function.zip"), nil),
		Handler: jsii.String("main"),
	})

	return stack
}

func main() {
	// defer - this keyword tells go code to run the particular line at the end after everything has completed execution
	// jsii - framework built by aws to allow us to use cdk with a language that is not typescript
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewGoAwsCdkStack(app, "GoAwsCdkStack", &GoAwsCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
