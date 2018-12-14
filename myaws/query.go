package main

import (
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type  AV   dynamodb.AttributeValue

func main() {
	//
	//
	//
	sess, err := session.NewSession(&aws.Config{
						Region: 	aws.String("ap-southeast-2"),
						Credentials: 	credentials.NewSharedCredentials("", "bob-developer"),
						   })
    	if err != nil {
        	fmt.Println("Error creating session:")
		log.Panic(err)
    	}
	//
	//  Define log destination
	//
        errf,err := os.OpenFile("aws.http.log",os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)			// (*File, error)
        if err != nil {
    	   log.Panic(err)
        }
	fmt.Fprintf(errf,"\n\n*****************^^^%s^^^^*****************\n\n",os.Args[0][strings.LastIndex(os.Args[0],"/")+1:])
        //
        // setup logger 
        //
        var loggerfunc func(args ...interface{})
        loggerfunc  =  func(args ...interface{}) {
    	fmt.Fprintln(errf, args...) 
        }
    	fmt.Println("loggerfunc:  ",loggerfunc)
        //
        //  create client-service. Apply extra session config.  NB  aws.LoggerFunc is an adapter function.  It adapts loggerfunc to a aws.Logger interface.
        // 
        svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody).WithLogger(aws.LoggerFunc(loggerfunc)) )
	//
	// Operation:  Query    requires a Key expression containing both partition and sort key ie.  year=? AND title=?
	//
	// Create a KeyConditionBuilder using:   KeyAnd( left, right KeyConditionBuilder) KeyConditionBuilder
	//                                  
	//                                where left is Partition key value: expression.Key("year").Equal(expression.Value("2013"))
	//                                     right is Sort key value:      expression.Key("title").Equal(expression.Value("The Big New Movie")
        //
	kcond := expression.Key("year").Equal(expression.Value(2015)).And(expression.KeyBeginsWith(expression.Key("title"),"T"))
 									// note 2015 is a number not string.   Value accepts interface{} - ANY TYPE
	fcond:= expression.LessThanEqual(expression.Name("info.rating"), expression.Value(5) )
	//
	// create expression from builder.  expression methods provides the struct values in the GetItemInput
	//
        expr,err:= expression.NewBuilder().WithKeyCondition(kcond).WithFilter(fcond).Build()

	fmt.Println("Condition :  ",*(expr.Condition()));
	for k,v := range expr.Names() {
		fmt.Printf("\n %s  %s",k,*v)
	}
	fmt.Printf("\n Condition : %s \n",expr.Values());

	input := &dynamodb.QueryInput{
    		KeyConditionExpression:    expr.KeyCondition(),		
		FilterExpression:          expr.Filter(),
  		ExpressionAttributeNames:  expr.Names(),		
  		ExpressionAttributeValues: expr.Values(),
		}
	input =input.SetTableName("Movies").SetReturnConsumedCapacity("TOTAL")	//.SetConsistentRead(true)
	//
	//
	// 
	out_,err:=svc.Query(input)
	if err!=nil {
		log.Print(err)
	}
	if len(out_.Items) > 0 {
	        fmt.Println(out_.Items)
	} else {
		fmt.Println("No data found..")
	}
}
	

