package main

import (
	"github.com/alejandroq/AWS-Go-Lambas-Behind-ALB/internal/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
