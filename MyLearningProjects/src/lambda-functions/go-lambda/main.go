package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	//"github.com/aws/aws-labbda-go"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process request ID %f", request.ID),
		Ok:      true,
	}, nil
}
