package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
	Region: aws.String("ap-southeast-2"),
	Credentials: credentials.NewSharedCredentials("", "bob-developer"),
	})									      	// overrides AWS_PROFILE environment variable
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	result,err := svc.ListTables(&dynamodb.ListTablesInput{});				// defaults to limit 100
	if err != nil {
	    if aerr, ok := err.(awserr.Error); ok {
	        switch aerr.Code() {
	        case dynamodb.ErrCodeInternalServerError:
	            fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
	        default:
	            fmt.Println(aerr.Error())
	        }
	    } else {
	        // Print the error, cast err to awserr.Error to get the Code and
	        // Message from an error.
	        fmt.Println(err.Error())
	    }
	    return
	}
	
	fmt.Println(result)
	
}
