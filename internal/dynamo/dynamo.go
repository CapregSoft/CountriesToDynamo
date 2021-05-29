package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDynamoClient() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)

	// Create DynamoDB client
	return dynamodb.New(sess)
}
