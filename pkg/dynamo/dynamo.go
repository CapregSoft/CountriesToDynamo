package dynamo

import (
	"github.com/CapregSoft/CountriesToDynamo/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetDynamoClient returns the client for the DynamoDb
func GetDynamoClient() *dynamodb.DynamoDB {
	c := config.GetConfig()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(c.Region)},
	)
	if err != nil {
		panic(err)
	}
	// Create DynamoDB client
	return dynamodb.New(sess)
}
