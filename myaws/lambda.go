package main

import (
        "context"
        "log"
	"fmt"
  "time"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
 "github.com/aws/aws-lambda-go/lambda"
	lda "github.com/aws/aws-sdk-go/service/lambda"
 	"github.com/aws/aws-sdk-go/aws/awserr"
)

type output_ struct {
        Answer string 
}

func LongRunningHandler(ctx context.Context) (output_, error) {

        deadline, _ := ctx.Deadline()
        fmt.Printf("deadline from context: %#v\n",deadline)
        deadline = deadline.Add(-100 * time.Millisecond)
        timeoutChannel := time.After(time.Until(deadline))


svc := lda.New(session.New())
//json_:=aws.String(`{ "What is your name?": "Jim", "How old are you?": 33 }x `)
//fmt.Println("Json: ",*json_)
input := &lda.InvokeInput{
    ClientContext:  aws.String("MyApp"),
    FunctionName:   aws.String("test-lambda-stack-3-TestFunction-Z4VCSGXF6KBQ"),
    //InvocationType: aws.String("Event"),  // default is synchronous
    LogType:        aws.String("Tail"),
    Payload:        []byte(nil) , //*json_),
 //   Qualifier:      aws.String("1"),
}
result, err := svc.Invoke(input)
if err == nil {
   fmt.Println("NO ERROR.....")
} else {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case lda.ErrCodeServiceException:
            fmt.Println(lda.ErrCodeServiceException, aerr.Error())
        case lda.ErrCodeResourceNotFoundException:
            fmt.Println(lda.ErrCodeResourceNotFoundException, aerr.Error())
        case lda.ErrCodeInvalidRequestContentException:
            fmt.Println(lda.ErrCodeInvalidRequestContentException, aerr.Error())
        case lda.ErrCodeRequestTooLargeException:
            fmt.Println(lda.ErrCodeRequestTooLargeException, aerr.Error())
        case lda.ErrCodeUnsupportedMediaTypeException:
            fmt.Println(lda.ErrCodeUnsupportedMediaTypeException, aerr.Error())
        case lda.ErrCodeTooManyRequestsException:
            fmt.Println(lda.ErrCodeTooManyRequestsException, aerr.Error())
        case lda.ErrCodeInvalidParameterValueException:
            fmt.Println(lda.ErrCodeInvalidParameterValueException, aerr.Error())
        case lda.ErrCodeEC2UnexpectedException:
            fmt.Println(lda.ErrCodeEC2UnexpectedException, aerr.Error())
        case lda.ErrCodeSubnetIPAddressLimitReachedException:
            fmt.Println(lda.ErrCodeSubnetIPAddressLimitReachedException, aerr.Error())
        case lda.ErrCodeENILimitReachedException:
            fmt.Println(lda.ErrCodeENILimitReachedException, aerr.Error())
        case lda.ErrCodeEC2ThrottledException:
            fmt.Println(lda.ErrCodeEC2ThrottledException, aerr.Error())
        case lda.ErrCodeEC2AccessDeniedException:
            fmt.Println(lda.ErrCodeEC2AccessDeniedException, aerr.Error())
        case lda.ErrCodeInvalidSubnetIDException:
            fmt.Println(lda.ErrCodeInvalidSubnetIDException, aerr.Error())
        case lda.ErrCodeInvalidSecurityGroupIDException:
            fmt.Println(lda.ErrCodeInvalidSecurityGroupIDException, aerr.Error())
        case lda.ErrCodeInvalidZipFileException:
            fmt.Println(lda.ErrCodeInvalidZipFileException, aerr.Error())
        case lda.ErrCodeKMSDisabledException:
            fmt.Println(lda.ErrCodeKMSDisabledException, aerr.Error())
        case lda.ErrCodeKMSInvalidStateException:
            fmt.Println(lda.ErrCodeKMSInvalidStateException, aerr.Error())
        case lda.ErrCodeKMSAccessDeniedException:
            fmt.Println(lda.ErrCodeKMSAccessDeniedException, aerr.Error())
        case lda.ErrCodeKMSNotFoundException:
            fmt.Println(lda.ErrCodeKMSNotFoundException, aerr.Error())
        case lda.ErrCodeInvalidRuntimeException:
            fmt.Println(lda.ErrCodeInvalidRuntimeException, aerr.Error())
        default:
            fmt.Println(aerr.Error())
        }
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println("Error:  ",err.Error())
 } else {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println("Error2: ",err.Error())
  }
}
fmt.Printf("input.Payload: [%s]\n",string(input.Payload))
fmt.Printf("result.Payload: %s\n",result.Payload)
fmt.Printf("result: %#v\n",result)
exchange := output_{}
err = json.Unmarshal(result.Payload,&exchange);
if err != nil {
   exchange=output_{Answer: "Ross is 60 years old"}
}
fmt.Println(exchange.Answer)

        for {

                select {

                case <- timeoutChannel:
                        return "Context Deadline almost reached", nil
                default:
                        log.Print("hello!")
                        time.Sleep(50 * time.Millisecond)
                        return exchange, nil
                }
        }
}

func main() {
        lambda.Start(LongRunningHandler)
}
