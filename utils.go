package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"os"
)

type EachConnection struct {
	ConnectionID string `json:"connectionID"`
}

func getEnvVar(key string) string {
	return os.Getenv(key)
}

var AccessKey = getEnvVar("ACCESSKEY")
var SecretKey = getEnvVar("SECRETKEY")
var SecretToken = getEnvVar("SECRETTOKEN")
var Region = getEnvVar("REGION")
var WebSocketEndpoint = getEnvVar("WEBSOCKETENDPOINT")

func NewDynamoDBSession() (*dynamodb.DynamoDB) {
	fmt.Println("Trying to obtain a DynamoDB session")
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKey, SecretKey, SecretToken),
	})
	fmt.Println("returning the dynamoDB session")
	return dynamodb.New(sess)
}

func NewAPIGatewaySession() *apigatewaymanagementapi.ApiGatewayManagementApi {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKey, SecretKey, SecretToken),
		Endpoint: aws.String(WebSocketEndpoint),
	})
	return apigatewaymanagementapi.New(sess)
}
