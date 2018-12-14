package main

import (
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type  AV   dynamodb.AttributeValue

func main() {
	//
	//  create a config, session and finally a client
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
	//  Data - while statistically declared here would come from an external source
	//
	var key_ = [][]string{[]string{"year","N","2015"}, []string{"title","S","The Big New Movie"} }
	//
	//  Create key - map[string]*AttributeValue
	//
	avMap_ := make(map[string]*dynamodb.AttributeValue)
	var av_  *dynamodb.AttributeValue
        for _,v := range key_ {
	    av_=new(dynamodb.AttributeValue)					// use new to allocate new struct - new === make (slice,map,chan)
	    //  av_=&(AV{})   							// literal allocation does not work when pointer required??
	    switch v[1] {							//  manually build an Map[String]*AV type used in Key
		case "N": 
			avMap_[v[0]] = av_.SetN(v[2]) 
		case "S": 
			avMap_[v[0]] = av_.SetS(v[2]) 
	    }	
   	}
	//fmt.Printf("\n Key:  %#v", avMap_);
	//
	//  add conditional clause...rating must be over 8
	//
 	cond := expression.GreaterThan(expression.Name("info.rating"), expression.Value(8))	
								//builder := expression.NewBuilder().WithCondition(cond)
								//expr,err := builder.Build()
	expr,err := expression.NewBuilder().WithCondition(cond).Build()
	if err != nil {
		log.Panic(err)
	}
	//
	//  Create Input struct
	//
	input := &dynamodb.DeleteItemInput{
	        Key:       			avMap_,					// must be of type: Map[string]*AttributeValue, alternatively use KeyBuilder()
		ConditionExpression:  		expr.Condition(),
		ExpressionAttributeNames:	expr.Names(),
		ExpressionAttributeValues:	expr.Values(),
	}
	input = input.SetReturnValues("ALL_OLD").SetTableName("Movies")
	//
	//  Delete Item
	//
	output_, err := svc.DeleteItem(input)
	if err != nil {
	        fmt.Println("Got error calling DeleteItem")
	        fmt.Println(err.Error())
	        return
	}
	//
	//  Examine response - output struct
	//
	fmt.Println("Len output.attributes  ",len(output_.Attributes))
	if len(output_.Attributes) != 0 {
	         fmt.Printf("\n\nDeleted %#s   ",output_.Attributes)
	} else {
		fmt.Println("\n\n No item found to delete..")
	}
}
