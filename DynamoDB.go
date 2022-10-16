package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/nullwulf/loggly"
)

func dynamodbInsert(insert CmpResponse, lgglyClient *loggly.ClientType) {

	lgglyClient.EchoSend("info", "\nAttempting dynamodb call.")

	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
		//EndPoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	// marshal the movie struct into an aws attribute value
	insertMap, err := dynamodbattribute.MarshalMap(insert)
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
		panic("Cannot marshal movie into AttributeValue map")
	}

	// create the api params
	params := &dynamodb.PutItemInput{
		TableName: aws.String(""),
		Item:      insertMap,
	}

	// put the item
	resp, err := db.PutItem(params)
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
		return
	}

	// print the response data
	lgglyClient.EchoSend("debug", "Pre response info output")
	lgglyClient.EchoSend("info", resp.String())
}
