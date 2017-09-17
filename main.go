package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

const (
	defaultFuncName  = "ian-golang-test"
	defaultAWSRegion = "us-east-1"
)

// Launch the lambda function
func main() {
	count := flag.Int("count", 1, "Number of functions to invoke")
	funcName := flag.String("func", defaultFuncName, "Lambda function to invoke")
	invokeType := flag.String("invoke", "Event", "Lambda function invocation type (Event/RequestResonse)")
	awsRegion := flag.String("region", defaultAWSRegion, "AWS Region")
	flag.Parse()
	log.Printf("Executing Lambda function %s %d times", *funcName, *count)

	sess := session.Must(session.NewSession())
	svc := lambda.New(sess, aws.NewConfig().WithRegion(*awsRegion))

	for i := 0; i < *count; i++ {
		resp, err := svc.Invoke(&lambda.InvokeInput{
			FunctionName:   funcName,
			InvocationType: invokeType,
			Payload:        []byte(fmt.Sprintf(`{"seq": %d}`, i+1)),
		})
		if err != nil {
			log.Printf("error launching function #%d: %v", i+1, err)
			continue
		}
		log.Printf("launched AWS Lambda Function #%d: status: %d", i+1, *resp.StatusCode)
	}
}
