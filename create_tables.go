package main

import (
	c "blogging/models"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func main() {

	region, check := os.LookupEnv(c.EnvRegion)

	if (!check) {
		region = c.DefaultRegion
	}

	endpoint, check := os.LookupEnv(c.EnvEndpoint)

	if (!check) {
		endpoint = c.DefaultEndpoint
	}

	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
		//Endpoint: aws.String("http://127.0.0.1:8000"),
		Endpoint: aws.String(endpoint),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	//AUTHORS
	// create the api params
	params := &dynamodb.CreateTableInput{
		TableName: aws.String(c.AuthorsTable),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String(c.Email), KeyType: aws.String("HASH")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String(c.Email), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(100),
		},
	}

	// create the table
	resp, err := db.CreateTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the response data
	fmt.Println(resp)

	//BLOG
	// create the api params
	//noinspection GoInvalidCompositeLiteral
	params = &dynamodb.CreateTableInput{
		TableName: aws.String(c.BlogsTable),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String(c.Title), KeyType: aws.String("HASH")},
			{AttributeName: aws.String(c.SpaceName), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String(c.Title), AttributeType: aws.String("S")},
			{AttributeName: aws.String(c.SpaceName), AttributeType: aws.String("S")},
			{AttributeName: aws.String(c.AuthorEmail), AttributeType: aws.String("S")},
			{AttributeName: aws.String(c.Approved), AttributeType: aws.String("S")},
		},
		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex{
			{
				IndexName: aws.String(c.LocalIndexAuthor),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String(c.Title), KeyType: aws.String("HASH")},
					{AttributeName: aws.String(c.AuthorEmail), KeyType: aws.String("RANGE")},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: nil,
					ProjectionType:   aws.String("ALL"),
				},
			},
			{
				IndexName: aws.String(c.LocalIndexApproved),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String("title"), KeyType: aws.String("HASH")},
					{AttributeName: aws.String("approved"), KeyType: aws.String("RANGE")},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: nil,
					ProjectionType:   aws.String("ALL"),
				},
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String(c.GlobalIndexSpaceName),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String(c.SpaceName), KeyType: aws.String("HASH")},
					{AttributeName: aws.String(c.AuthorEmail), KeyType: aws.String("RANGE")},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: nil,
					ProjectionType:   aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(100),
				},
			},
			{
				IndexName: aws.String(c.GlobalIndexAuthorEmail),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String(c.AuthorEmail), KeyType: aws.String("HASH")},
					{AttributeName: aws.String(c.Title), KeyType: aws.String("RANGE")},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: nil,
					ProjectionType:   aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(100),
				},
			},
			{
				IndexName: aws.String(c.GlobalIndexApproved),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String(c.Approved), KeyType: aws.String("HASH")},
					{AttributeName: aws.String(c.Title), KeyType: aws.String("RANGE")},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: nil,
					ProjectionType:   aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(100),
				},
			},
			{
				IndexName: aws.String(c.GlobalIndexSpaceNameApproved),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String(c.Approved), KeyType: aws.String("HASH")},
					{AttributeName: aws.String(c.SpaceName), KeyType: aws.String("RANGE")},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: nil,
					ProjectionType:   aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(100),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(100),
		},
	}

	// create the table
	resp, err = db.CreateTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the response data
	fmt.Println(resp)

	//Space
	// create the api params
	params = &dynamodb.CreateTableInput{
		TableName: aws.String(c.SpacesTable),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String(c.SpaceName), KeyType: aws.String("HASH")},
			{AttributeName: aws.String(c.OwnerEmail), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String(c.SpaceName), AttributeType: aws.String("S")},
			{AttributeName: aws.String(c.OwnerEmail), AttributeType: aws.String("S")},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String(c.GlobalIndexOwnerEmail),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String(c.OwnerEmail), KeyType: aws.String("HASH")},
					{AttributeName: aws.String(c.SpaceName), KeyType: aws.String("RANGE")},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: nil,
					ProjectionType:   aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(100),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(100),
		},
	}

	// create the table
	resp, err = db.CreateTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print the response data
	fmt.Println(resp)
}
