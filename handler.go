package main

import (
	"log"
	"time"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// Handle is what gets run on AWS Lambda. This function just executes a sleep
// to allow a large number of functions to stack up, to exercise limits of
// maximum parallel invocation.
func Handle(evt interface{}, ctx *runtime.Context) (map[string]interface{}, error) {
	log.Printf(
		"Lambda function %s invoked with request %s at %s",
		ctx.FunctionName,
		ctx.AWSRequestID,
		time.Now())

	// Sleep for a minute to allow functions to stack up.
	time.Sleep(60 * time.Second)

	log.Printf(
		"Lambda function %s finished for request %s with %d milliseconds remaining",
		ctx.FunctionName,
		ctx.AWSRequestID,
		ctx.RemainingTimeInMillis())

	return map[string]interface{}{
		"AWSRequestID":          ctx.AWSRequestID,
		"TimeRemainingInMillis": ctx.RemainingTimeInMillis(),
		"CurrentTime":           time.Now().String(),
	}, nil
}
