package main

import (
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
	Region: aws.String("ap-southeast-2"),
	Credentials: credentials.NewSharedCredentials("", "bob-developer"),
	})
    	if err != nil {
        	fmt.Println("Error creating session:")
		log.Panic(err)
    	}
    	// Create DynamoDB client
    	svc := dynamodb.New(sess)


    input := &dynamodb.DeleteItemInput{
        Key: map[string]*dynamodb.AttributeValue{
            "year": {
                N: aws.String("2015"),
            },
            "title": {
                S: aws.String("The Big New Movie"),
            },
        },
        TableName: aws.String("Movies"),
    }

    _, err = svc.DeleteItem(input)

    if err != nil {
        fmt.Println("Got error calling DeleteItem")
        fmt.Println(err.Error())
        return
    }

    fmt.Println("Deleted 'The Big New Movie' (2015)")
}
