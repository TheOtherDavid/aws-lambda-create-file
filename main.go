package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func handleRequest() (string, error) {
	path := os.Getenv("FILE_PATH")
	filename := path + "newFile.csv"
	createFile(filename)
	return "Success", nil
}

func main() {
	lambda.Start(handleRequest)
}

func createFile(filename string) {
	_, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
