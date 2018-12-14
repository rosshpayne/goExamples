package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "log"
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Create structs to hold info about new item
type ItemInfo struct {
    Plot string			`json:"plot"`
    Rating float64		`json:"rating"`
}

type Item struct {
    Year int			`json:"year"`
    Title string		`json:"title"`
    Info ItemInfo		`json:"info"`			// embedded struct - missing some atributes compared to data in first item
}

// Get table items from JSON file
func getItems() []Item {
    raw, err := ioutil.ReadFile("./movie_data.json")

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var items []Item						// decode destination. Slice of struct maps to an array of json objects.
    err=json.Unmarshal(raw, &items)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	fmt.Println(items)
    return items
}

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
        // Get table items from movie_data.json
        items := getItems()
    
        // Add each item to Movies table:
        for _, item := range items {
    
            av, err := dynamodbattribute.MarshalMap(item)			// encode go struct item Dynamodb attributevalue (AV) -> Map[string]*AV
    	    if err != nil {
    	      log.Fatal(err)
    	    }
    
    	for k,v := range av {						//   AV can have only one  field set
    	   fmt.Println("*****M key : ",k)
    	   fmt.Println("      			AV.M  : ",v)
    	   if v.S != nil {
    		fmt.Println("                        AV.S  : ",*(v.S))	
     	   } else {
    	        v=v.SetS("AutoSet string..")
    		fmt.Println("                        AV.S  : ",*(v.S))	
    	   }
    	   fmt.Println("      			AV.SS  : ",v.SS)
    	   if v.N != nil {
    	      fmt.Println("      			AV.N  : ",*(v.N))
    	   } else {
    	      fmt.Println("      			AV.NS : ",v.NS)
    	      fmt.Println("      			AV.N  : ",v.N)
    	   }
           if err != nil {
                fmt.Println("Got error marshalling map:")
                fmt.Println(err.Error())
                os.Exit(1)
           }
    
           input := &dynamodb.PutItemInput{				
                Item: av,							// takes Map[string]*AV from Marshal and loads into Table 
                TableName: aws.String("Movies"),
           }
           _, err = svc.PutItem(input)					// insert item & create attributes as required.
    
           if err != nil {
                fmt.Println("Got error calling PutItem:")
                fmt.Println(err.Error())
                os.Exit(1)
           }
    
           fmt.Println("Successfully added '",item.Title,"' (",item.Year,") to Movies table")
    
       }
}
