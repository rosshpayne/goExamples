package main

import (
    "fmt"
    "log"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

type  AV   dynamodb.AttributeValue

func main() {
	//
	//
	//
	pk:=os.Args[1]
	sortkey:=os.Args[2]
	
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
	fmt.Fprint(errf,"\n\n*****************^^^^^^^^^*****************\n\n")
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
	//  Data - while statistically declared here would come from an external source
	//
	var key_ = [][]string{[]string{"year","N",pk}, []string{"title","S",sortkey} }
	//
	//  Create key map[string]*AttributeValue:  Note there is no expression builder that generates a key of this data type.
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
	//  Create *Input struct
	//
	input := &dynamodb.GetItemInput{TableName: aws.String("Movies")}
	//
	// Alternatively:  use Set* methods to assign Input fields
	//
	input = input.SetKey(avMap_).SetReturnConsumedCapacity("TOTAL")
	//
	//  Delete Item
	//
	output_, err := svc.GetItem(input)
	if err != nil {
	        log.Print("Got error calling GetItem")
	        log.Panic(err)
	}
	//
	//  Examine response - output struct
	//
	fmt.Println("Len output.Item  ",len(output_.Item))
	if len(output_.Item) == 0 {
		fmt.Println("\n\n No item found ")
	} else {
		fmt.Println(output_.Item)
		fmt.Println(output_.ConsumedCapacity)
	}
}
