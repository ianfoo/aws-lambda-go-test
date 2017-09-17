# AWS Lambda Go Test

This project uses [aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim)
to create an AWS Lambda function written in Go.

## Building

Run `make`.

## Deploying

You'll need to create a new Python 2.7 AWS Lambda function, and upload the zip file
produced by the build, and set the handler function to `handler.Handle`.

## Running

Usage:
```
  -count int
    	Number of functions to invoke (default 1)
  -func string
    	Lambda function to invoke (default "ian-golang-test")
  -invoke string
    	Lambda function invocation type (Event/RequestResonse) (default "Event")
  -region string
    	AWS Region (default "us-east-1")
```
Set `<function_name>` to match the name of the function you created in AWS Lambda.

