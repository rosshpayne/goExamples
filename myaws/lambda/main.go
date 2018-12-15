package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	lda "github.com/aws/aws-sdk-go/service/lambda"
	_ "log"
	"time"
)

type output_ struct {
	Answer string
}

type input_ struct {
	YourName string `json:"What is your name?"`
	Age      int    `json:"How old are you?"`
}

//func LongRunningHandler(ctx context.Context) (output_, error) {
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	deadline, _ := ctx.Deadline()
	fmt.Printf("deadline from context: %#v\n", deadline)
	deadline = deadline.Add(-100 * time.Millisecond)
	//CChannel := time.After(time.Until(deadline))

	config := aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := lda.New(sess)
	//svc := lda.New(session.NewSession(&aws.Config{Region:, Loglevel: aws.LogDebugWithHTTPBody}))

	ctxInput := `{ "NameX": "ValueY" } `
	b64encoded := base64.StdEncoding.EncodeToString([]byte(ctxInput))
	b64json_ := aws.String(b64encoded)

	myInput := input_{"Rossco", 21}
	payload_, err := json.Marshal(&myInput)
	if err != nil {
		panic(err)
	}
	input := &lda.InvokeInput{
		ClientContext: b64json_, // internally base64 encoded and attached to header: X-Amz-Client-Context: eyAiTmFtZVgiOiAiVmFsdWVZIiB9IA==
		FunctionName:  aws.String("test-lambda-stack-3-TestFunction-Z4VCSGXF6KBQ"),
		//InvocationType: aws.String("Event"),  // default is synchronous
		LogType: aws.String("Tail"),
		Payload: payload_,
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
			fmt.Println("Error:  ", err.Error())
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println("Error2: ", err.Error())
		}
	}
	fmt.Printf("input.Payload: [%s]\n", string(input.Payload))
	fmt.Printf("result.Payload: %s\n", result.Payload)
	fmt.Printf("result: %#v\n", result)
	exchange := output_{}
	err = json.Unmarshal(result.Payload, &exchange)
	if err != nil {
		panic(err)
	}
	//exchange = output_{Answer: "Ross is 70 years old"}
	fmt.Println("Exchange: ", exchange.Answer)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       exchange.Answer,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
