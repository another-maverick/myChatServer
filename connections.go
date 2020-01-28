package main

import(
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
	lambda.Start(ConnectionHandler)
}

func ConnectionHandler(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	eachConnection := EachConnection{ConnectionID: request.RequestContext.ConnectionID,}
	connValue, _ := dynamodbattribute.MarshalMap(eachConnection)
	item := &dynamodb.PutItemInput{
		Item:      connValue,
		TableName: aws.String("myChatConnection"),
	}
	fmt.Sprintf("here is the connection ID and value dynamoDB item %v %v \n",  eachConnection.ConnectionID, connValue )
	dynamodbSession := NewDynamoDBSession()
	fmt.Println("got new dynamo DB session from utils...")
	dynamodbSession.PutItem(item)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil
}
