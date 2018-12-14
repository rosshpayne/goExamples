package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/request"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "fmt"
    "log"
    "os"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
	Region: 	aws.String("ap-southeast-2"),
	Credentials: 	credentials.NewSharedCredentials("", "bob-developer"),
            	 })
    	if err != nil {
        	fmt.Println("Error creating session:")
		log.Panic(err)
    	}
	fmt.Println("Validate:        ",sess.Handlers.Validate)
	fmt.Println("Build:           ",sess.Handlers.Build)
	fmt.Println("Sign             ",sess.Handlers.Sign)								// HandleirList is a slice of HandlerFuncs 
	fmt.Println("Send             ",sess.Handlers.Send)
	fmt.Println("ValidateResp     ",sess.Handlers.ValidateResponse)
	fmt.Println("Unmarshal        ",sess.Handlers.Unmarshal)
	fmt.Println("UnmrashalMeta    ",sess.Handlers.UnmarshalMeta)
	fmt.Println("UnmarshalError   ",sess.Handlers.UnmarshalError)
	fmt.Println("Retry            ",sess.Handlers.Retry)
	fmt.Println("AfterRetry       ",sess.Handlers.AfterRetry)
        fmt.Printf("LogLevel LogDebugWithSigning         %b\n",aws.LogDebugWithSigning)
        fmt.Printf("LogLevel LogDebugWithHTTPBody        %b\n",aws.LogDebugWithHTTPBody)
        fmt.Printf("LogLevel LogDebugWithRequestRetries  %b\n",aws.LogDebugWithRequestRetries)
        fmt.Printf("LogLevel LogDebugWithRequestErrors   %b\n",aws.LogDebugWithRequestErrors)
        //
        // Create DynamoDB client and expose HTTP requests/responses
        //
        errf,err := os.OpenFile("aws.http.log",os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)			// (*File, error)
        if err != nil {
    	   log.Panic(err)
        }
	fmt.Fprint(errf,"\n\n*****************^^^^^^^^^*****************\n\n")
        //
        // setup logger to file
        //
        var loggerfunc func(args ...interface{})
        loggerfunc  =  func(args ...interface{}) {
    	fmt.Fprintln(errf, args...) 
        }
    	fmt.Println("loggerfunc:  ",loggerfunc)
        //
        //  create client-service. Apply extra session config.  NB  aws.LoggerFunc is an adapter function.  It adapts loggerfunc to a aws.Logger interface.
	//    New returns DynamoDB type
	//    type DynamoDB struct { 					//an embedded anonymous type (ptr to Client) which promotes its fields and methods to svc.
	//              *client.Client 
	//    } 
	//    type Client struct {
    	//		request.Retryer
    	//		metadata.ClientInfo
	//		
    	//		Config   aws.Config				// Config & Handlers populated from sess (type Session) ??
    	//		Handlers request.Handlers               	//   dynamoDB.New gives opportunity to override (?? add+update) Config in second argument
	//    }								//   which can be nil so config is defined soley from sess.
        //	with methods:  AddDebugHandlers(..), NewRequest(..)
	//     DynamoDB type has lots & lots of dynamodb methods e.g. ListTables below. 
	//     so svc has methods promoted from its fields (client) as well as methods explicitly written on it.
	//       each method has access to the receiver ie. the pointer to the DynameDB struct containing Client struct fields and methods
	// 
        svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody).WithLogger(aws.LoggerFunc(loggerfunc)) )
	//
        //  alternatively, if no change to Config. Second argument is verardic meaning 0..n arguments can be added. In this case its really just 0 or 1.
        //      svc := dynamodb.New(sess)
        //
        // Add "CustomHeader" header with value of 10.  NB:  svc is instance of DynamoDB struct which has a single anonymous embedded field of Client{.. aws.Config, aws.Handlers }
        //    Send is field of Handlers. Send is a HandleList which has numerous methods ie. PushFront which adds func to Send handler list.
        //
        svc.Handlers.Send.PushFront(func(r *request.Request) {
            r.HTTPRequest.Header.Set("CustomHeader", fmt.Sprintf("%d", 10))
        })
        //fmt.Printf("Send HandlerList:  %#v",svc.Handlers.Send)
        // Call ListTables just to see HTTP request/response
        // The request should have the CustomHeader set to 10
    
        _, _ = svc.ListTables(&dynamodb.ListTablesInput{})
}
