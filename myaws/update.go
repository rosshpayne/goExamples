package main

import (
 	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	da "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	ddb "github.com/aws/aws-sdk-go/service/dynamodb"
	)

func main () {

	sess, err := session.NewSession(&aws.Config{
    	Region: aws.String("ap-southeast-2")},
	)

	// Create DynamoDB client
	svc := ddb.New(sess)


        input := &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String("0.5"),
            },
        },
        TableName: aws.String("Customer"),
        Key: map[string]*dynamodb.AttributeValue{
            "year": {
                N: aws.String("2015"),
            },
            "title": {
                S: aws.String("The Big New Movie"),
            },
        },
        ReturnValues:     aws.String("ALL_NEW"),
        UpdateExpression: aws.String("set info.rating = :r"),
        }

    _, err = svc.UpdateItem(input)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("Successfully updated 'The Big New Movie' (2015) rating to 0.5")
}
