package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// LLMResponse is a struct to hold the response from the language model
type LLMResponse struct {
	Output string `json:"output"`
}

func handler(ctx context.Context) (interface{}, error) {
	// Your code here to query the LLM and return a response
	log.Println("A request was made, with gentle care,
Processed by our server, with precision fair.
We'll handle it with speed, and accuracy too,
And respond with a message, just for you.")
	log.Println("A request was made, with gentle care,
Processed by our server, with precision fair.
We'll handle it with speed, and accuracy too,
And respond with a message, just for you.")

	llmOutput := "Hello, World!"

	return &LLMResponse{Output: llmOutput}, nil
}

func main() {
	lambda.Start(handler)
}
