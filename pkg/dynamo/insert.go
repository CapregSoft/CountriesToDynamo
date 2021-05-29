package dynamo

import (
	"fmt"
	"log"

	"github.com/CapregSoft/CountriesToDynamo/pkg/config"
	"github.com/CapregSoft/CountriesToDynamo/pkg/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func Insert(records []model.Country) {
	config := config.GetConfig()
	dynamo := GetDynamoClient()

	for _, record := range records {
		log.Printf("Inserting %v", record.CountryName)
		item, err := dynamodbattribute.MarshalMap(record)
		if err != nil {
			panic(fmt.Sprintf("Failed to marshal country record, %v", err))
		}

		req, _ := dynamo.PutItemRequest(&dynamodb.PutItemInput{
			TableName: aws.String(config.TableName),
			Item:      item},
		)
		err = req.Send()
		if err != nil {
			panic(err)
		}
	}
}
