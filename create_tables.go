package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-west-1"),
		Endpoint: aws.String("http://127.0.0.1:8000"),
		//EndPoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	//AUTHORS
	// create the api params
	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Authors"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("email"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("display_name"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("email"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("display_name"), AttributeType: aws.String("S")},
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
		TableName: aws.String("Blogs"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("title"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("space_name"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("title"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("space_name"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("author_name"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("approved"), AttributeType: aws.String("S")},
		},
		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex {
			{
				IndexName: aws.String("local_index_author"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String("title"), KeyType: aws.String("HASH")},
					{AttributeName: aws.String("author_name"), KeyType: aws.String("RANGE")},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: nil,
					ProjectionType:   aws.String("ALL"),
				},
			},
			{
				IndexName: aws.String("local_index_approved"),
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
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex {
			{
				IndexName: aws.String("global_index_space_name"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String("space_name"), KeyType: aws.String("HASH")},
					{AttributeName: aws.String("author_name"), KeyType: aws.String("RANGE")},
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
				IndexName: aws.String("global_index_author_name"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String("author_name"), KeyType: aws.String("HASH")},
					{AttributeName: aws.String("title"), KeyType: aws.String("RANGE")},
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
				IndexName: aws.String("global_index_approved"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{AttributeName: aws.String("approved"), KeyType: aws.String("HASH")},
					{AttributeName: aws.String("title"), KeyType: aws.String("RANGE")},
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

	/*//Space
	// create the api params
	params = &dynamodb.CreateTableInput{
		TableName: aws.String("Spaces"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("space_name"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("owner_name"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("space_name"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("owner_name"), AttributeType: aws.String("S")},
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
*/
	//Approval
	// create the api params
	/*params = &dynamodb.CreateTableInput{
		TableName: aws.String("Approvals"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("approval_id"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("space_name"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("approval_id"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("space_name"), AttributeType: aws.String("S")},
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
	}*/

	// print the response data
	fmt.Println(resp)
}