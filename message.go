package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	lambda.Start(HandleMessage)
}

// RequestPayload represents the request body sent by the Socket
type RequestData struct {
	Message string `json:"message"`
}

type MessageContent struct {
	Message      string `json:"message"`
	ConnectionID string `json:"connectionId"`
}

func HandleMessage(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse the request
	var requestData RequestData
	json.Unmarshal([]byte(request.Body), &requestData)

	dynamodbSession := NewDynamoDBSession()

	// Read the chat connection table
	input := &dynamodb.ScanInput{
		TableName: aws.String("myChatConnection"),
	}
	scan, _ := dynamodbSession.Scan(input)

	// Parse the table data in the output variable
	var output []EachConnection
	dynamodbattribute.UnmarshalListOfMaps(scan.Items, &output)

	apigatewaySession := NewAPIGatewaySession()

	// Encode the message data with Message and Connection ID
	messageContent := &MessageContent{
		Message:      requestData.Message,
		ConnectionID: request.RequestContext.ConnectionID,
	}
	jsonData, _ := json.Marshal(messageContent)

	// Send the message for each connection ID
	for _, item := range output {
		connectionInput := &apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(item.ConnectionID),
			Data:         jsonData,
		}
		_, err := apigatewaySession.PostToConnection(connectionInput)
		if err != nil {
			fmt.Println(err)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil
}