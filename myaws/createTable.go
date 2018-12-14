package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    _ "github.com/aws/aws-sdk-go/aws/awserr"
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

	// Create table Movies
	input := &dynamodb.CreateTableInput{
	     AttributeDefinitions: []*dynamodb.AttributeDefinition{			// table DDL is specified with dynamo.AttributeDefinition(s) 
	         {									// the data is specified  withe dynamo.AttributeValue(s) which is marshaled  
	             AttributeName: aws.String("year"),					// from Go types
	             AttributeType: aws.String("N"),
	         },
	         {
	             AttributeName: aws.String("title"),
	             AttributeType: aws.String("S"),
	         },
	     },
	     KeySchema: []*dynamodb.KeySchemaElement{
	         {
	             AttributeName: aws.String("year"),
	             KeyType:       aws.String("HASH"),
	         },
	         {
	             AttributeName: aws.String("title"),
	             KeyType:       aws.String("RANGE"),
	         },
	     },
	     ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
	         ReadCapacityUnits:  aws.Int64(5),
	         WriteCapacityUnits: aws.Int64(5),
	     },
	     TableName: aws.String("Movies"),
	}
	
	_, err = svc.CreateTable(input)
	
	if err != nil {
	     fmt.Println("Got error calling CreateTable:")
	     fmt.Println(err.Error())
	     os.Exit(1)
	}
	
	fmt.Println("Created the table Movies in us-west-2")
	}
