package main

import (
        "fmt"
        "log"
        "os"

        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        _ "github.com/aws/aws-sdk-go/aws/credentials"
        _ "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
        "github.com/aws/aws-sdk-go/service/dynamodb"
        )

func main () {

        // default credentials using default credentials chain
        //  1. In system environment variables: AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY.
        //  2.  In the Go system properties: aws.accessKeyId and aws.secretKey. ???
        //  3.  In the default credentials file (the location of this file varies by platform).
        //  4.  In the instance profile credentials, which exist within the instance metadata associated 
        //      with the IAM role for the EC2 instance.
        sess, err := session.NewSession(&aws.Config{
                                        				Region: aws.String("ap-southeast-2"),
                                       				 })
        if err != nil {
                fmt.Println("Error creating session:")
                log.Panic(err)
        }
        // Create DynamoDB client
        svc := dynamodb.New(sess)

        tablenameinput:= &dynamodb.DescribeTableInput{}

        for _,tab := range os.Args[1:] {
                output_,err := svc.DescribeTable( tablenameinput.SetTableName(tab))
                if err != nil {
                        log.Panic(err)
                }
                fmt.Println(output_.GoString())
        }
}
