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
    "time"
    "strings"
    "context"
    "sync"
    "strconv"
)

func main() {

	fmt.Printf("\nArg %v %[1]T\n",os.Args[1])
        seconds_,err:=strconv.Atoi(os.Args[1])
        if err!=nil {
	    log.Fatal(err)
        }
	fmt.Printf("\nArg %d\n",seconds_)

        cancelIn,_:= time.ParseDuration(strconv.Itoa(seconds_)+"s")
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
        // Create DynamoDB client and expose HTTP requests/responses
        //
        errf,err := os.OpenFile("aws.http.log",os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)			// (*File, error)
        if err != nil {
    	log.Panic(err)
        }
 	fmt.Fprintf(errf,"\n\n*****************^^^%s^^^^*****************\n\n",os.Args[0][strings.LastIndex(os.Args[0],"/")+1:])
	if err != nil {
	   log.Panic(err)
	}
        //
        // setup logger to file
        //
        var loggerfunc func(args ...interface{})
        loggerfunc  =  func(args ...interface{}) { fmt.Fprintln(errf, args...) }
        //
        //  create client-service. Apply extra session config.  NB  aws.LoggerFunc is an adapter function.  It adapts loggerfunc to a aws.Logger interface.
        // 
        svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody).WithLogger(aws.LoggerFunc(loggerfunc)) )
        //
        // Add "CustomHeader" header with value of 10.  NB:  svc is instance of DynamoDB struct which has a single anonymous embedded field of Client{.. aws.Config, aws.Handlers }
        //    Send is field of Handlers. Send is a HandleList which has numerous methods ie. PushFront which adds func to Send handler list.
        //
        svc.Handlers.Send.PushFront(func(r *request.Request) {
            r.HTTPRequest.Header.Set("CustomHeader", fmt.Sprintf("%d", 10))
        })
	//
        // listtables with context.
        // 
	var wg sync.WaitGroup
	wg.Add(1)
	
   	ctx,cancelCtx:=context.WithCancel(context.Background())
	defer cancelCtx()							// always cancel to prevent leak of context. Closes channel for one thing.
        go func() {defer wg.Done()
		   time.Sleep(2*time.Second)   				// time.Second is a duration type
	           outp,err:=svc.ListTablesWithContext(ctx,&dynamodb.ListTablesInput{}) // no Option(s)... see request.Option
		   if err != nil {
		      log.Print(err)
		   }
		   log.Println(outp)
		   fmt.Println("Exit GoRoutine..")
        }()
	if err != nil {
	    log.Print(err)
	} else {
	    start := time.Now()
	    time.AfterFunc(cancelIn ,func (){cancelCtx()})			// runs in its own GoRoutine.  After CancelIn seconds function is executed.
	    t := time.Now()
	    elapsed := t.Sub(start)
	    fmt.Println(elapsed) 
	}
	wg.Wait()
	//
}
