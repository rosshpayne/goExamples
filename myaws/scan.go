
package main

import (
    "fmt"
    "os"
    "log" 

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// Create structs to hold info about new item
type ItemInfo struct {
    Plot string`json:"plot"`
    Rating float64`json:"rating"`
}

type Item struct {
    Year int`json:"year"`
    Title string`json:"title"`
    Info ItemInfo`json:"info"`
}

// Get the movies with a minimum rating of 8.0 in 2011
func main() {
	//
	//   query data
	//
	min_rating := 8.0
    	year := 2015
	//
	//   config, session & client setup
	//
	sess, err := session.NewSession(&aws.Config{
	Region: aws.String("ap-southeast-2"),
	Credentials: credentials.NewSharedCredentials("", "bob-developer"),
	})
    	if err != nil {
        	fmt.Println("Error creating session:")
		log.Panic(err)
    	}
    	svc := dynamodb.New(sess)
	//
        // Create filter expression
	//
        filt := expression.Name("year").Equal(expression.Value(year))
    
        // Or we could get by ratings and pull out those with the right year later
        //    filt := expression.Name("info.rating").LessThan(expression.Value(min_rating))
    
        // Get back the title, year, and rating
        proj := expression.NamesList(expression.Name("title"), expression.Name("year"), expression.Name("info.rating"))
    
        expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
    
        if err != nil {
            fmt.Println("Got error building expression:")
            fmt.Println(err.Error())
            os.Exit(1)
        }
    
        // Build the query input parameters
        params := &dynamodb.ScanInput{
            ExpressionAttributeNames:  expr.Names(),
            ExpressionAttributeValues: expr.Values(),
            FilterExpression:          expr.Filter(),
            ProjectionExpression:      expr.Projection(),
            TableName:                 aws.String("Movies"),
        }
    
        result, err := svc.Scan(params)
    
        if err != nil {
            fmt.Println("Query API call failed:")
            fmt.Println((err.Error()))
            os.Exit(1)
        }
    
        num_items := 0
    
        for _, i := range result.Items {
            item := Item{}
    
            err = dynamodbattribute.UnmarshalMap(i, &item)
    
            if err != nil {
                fmt.Println("Got error unmarshalling:")
                fmt.Println(err.Error())
                os.Exit(1)
            }
    
            // Which ones had a higher rating?
            if item.Info.Rating < min_rating {
                // Or it we had filtered by rating previously:
                //   if item.Year == year {
                num_items += 1
    
                fmt.Println("Title: ", item.Title)
                fmt.Println("Rating:", item.Info.Rating)
                fmt.Println()
            }
        }
    
        fmt.Println("Found", num_items, "movie(s) with a rating below", min_rating, "in", year)
}
