package db

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/kmathew/blogging/models"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(models.DefaultRegion))

//var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-west-1").WithEndpoint("http://127.0.0.1:8000"))

func RegisterAuthor(name string, displayName string, email string) (string, error) {
	author := models.Author{
		Name:        name,
		DisplayName: displayName,
		Email:       email,
	}

	// marshal the movie struct into an aws attribute value
	authorAVMap, err := dynamodbattribute.MarshalMap(author)
	if err != nil {
		panic("Cannot marshal author into AttributeValue map")
	}

	// create the api params
	params := &dynamodb.PutItemInput{
		TableName: aws.String(models.AuthorsTable),
		Item:      authorAVMap,
	}

	// put the item
	resp, err := db.PutItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return "", err
	}

	return resp.GoString(), nil
}

func CreateSpace(name string, ownerEmail string) error {

	space := models.Space{
		Name:       name,
		OwnerEmail: ownerEmail,
	}

	// marshal the movie struct into an aws attribute value
	spaceAVMap, err := dynamodbattribute.MarshalMap(space)
	if err != nil {
		panic("Cannot marshal space into AttributeValue map")
	}

	// create the api params
	params := &dynamodb.PutItemInput{
		TableName: aws.String(models.SpacesTable),
		Item:      spaceAVMap,
	}

	// put the item
	resp, err := db.PutItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return err
	}

	// print the response data
	fmt.Println(resp)

	return nil
}

func GetAuthor(email string) (models.Author, error) {
	var author models.Author
	params := &dynamodb.GetItemInput{
		TableName: aws.String(models.AuthorsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
	}

	// read the item
	resp, err := db.GetItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return author, err
	}

	// dump the response data
	fmt.Println(resp)

	// unmarshal the dynamodb attribute values into a custom struct

	err = dynamodbattribute.UnmarshalMap(resp.Item, &author)

	return author, nil
}

func GetAuthorSpace(authorEmail string) ([]models.Space, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(authorEmail),
			},
		},
		KeyConditionExpression: aws.String("owner_email = :v1"),
		ProjectionExpression:   aws.String(models.SpaceName),
		TableName:              aws.String(models.SpacesTable),
		IndexName:              aws.String(models.GlobalIndexOwnerEmail),
	}

	result, err := db.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	spaces := []models.Space{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &spaces)

	return spaces, nil
}

func CreateBlog(title string, content []byte, spaceName string, authorEmail string) (string, error) {
	//generate id
	//insert blog into table
	//if spaceid belongs to authorid.... then auto approve
	//if not then send blog for approval

	blog := models.Blog{
		Title:       title,
		Content:     content,
		SpaceName:   spaceName,
		AuthorEmail: authorEmail,
		Approved:    models.False,
	}

	//author, err := get
	// marshal the movie struct into an aws attribute value
	blogAVMap, err := dynamodbattribute.MarshalMap(blog)
	if err != nil {
		panic("Cannot marshal space into AttributeValue map")
	}

	// create the api params
	params := &dynamodb.PutItemInput{
		TableName: aws.String("Blogs"),
		Item:      blogAVMap,
	}

	// put the item
	resp, err := db.PutItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return "", err
	}

	return resp.GoString(), nil
}

func GetBlog(title string) ([]models.Blog, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(title),
			},
		},
		KeyConditionExpression: aws.String("title = :v1"),
		TableName:              aws.String(models.BlogsTable),
	}

	blogs := []models.Blog{}

	result, err := db.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &blogs)

	return blogs, nil
}

func ApproveBlog(spaceName string, blogTitle string, approverEmail string) (string, error) {
	blogsList, err := getBlog(blogTitle)
	if err != nil {
		return "", err
	}
	spacesList, err := getAuthorSpace(approverEmail)

	if err != nil {
		return "", err
	}

	if blogsList[0].SpaceName == spaceName && spacesList[0].Name == spaceName {
		//means we can approve it
		input := &dynamodb.UpdateItemInput{
			ExpressionAttributeNames: map[string]*string{
				"#AT": aws.String(models.Approved),
			},
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":t": {
					S: aws.String(models.True),
				},
			},
			Key: map[string]*dynamodb.AttributeValue{
				models.Title: {
					S: aws.String(blogTitle),
				},
				models.SpaceName: {
					S: aws.String(spaceName),
				},
			},
			ReturnValues:     aws.String("ALL_NEW"),
			TableName:        aws.String(models.BlogsTable),
			UpdateExpression: aws.String("SET #AT = :t"),
		}
		result, err := db.UpdateItem(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeConditionalCheckFailedException:
					fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
				case dynamodb.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case dynamodb.ErrCodeResourceNotFoundException:
					fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
				case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
					fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
				case dynamodb.ErrCodeTransactionConflictException:
					fmt.Println(dynamodb.ErrCodeTransactionConflictException, aerr.Error())
				case dynamodb.ErrCodeRequestLimitExceeded:
					fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return "", err
		}
		return result.GoString(), nil
	}
	//fail.. cant approve it because blog/space doesnt exist or not owner of space
	return "", errors.New("fail.. cant approve it because blog/space doesnt exist or not owner of space")

}

func GetBlogsForSpaceName(spaceName string) ([]models.Blog, error) {
	//return all blogs associated with spaceid that are approved
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(spaceName),
			},
			":v2": {
				S: aws.String(models.True),
			},
		},
		KeyConditionExpression: aws.String("space_name = :v1 and approved = :v2"),
		TableName:              aws.String(models.BlogsTable),
		IndexName:              aws.String(models.GlobalIndexSpaceName),
	}

	result, err := db.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	blogs := []models.Blog{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &blogs)
	return blogs, nil
}

func GetApprovedBlogs() ([]models.Blog, error) {
	//get unapproved blogs
	//return all blogs associated with spaceid that are unapproved
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(models.True),
			},
		},
		KeyConditionExpression: aws.String("approved = :v1"),
		TableName:              aws.String(models.BlogsTable),
		IndexName:              aws.String(models.GlobalIndexApproved),
	}

	result, err := db.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	blogs := []models.Blog{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &blogs)
	return blogs, nil
}

func GetAllUnapprovedBlogs() ([]models.Blog, error) {
	//get unapproved blogs
	//return all blogs associated with spaceid that are unapproved
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(models.False),
			},
		},
		KeyConditionExpression: aws.String("approved = :v1"),
		TableName:              aws.String(models.BlogsTable),
		IndexName:              aws.String(models.GlobalIndexApproved),
	}

	result, err := db.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	blogs := []models.Blog{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &blogs)
	return blogs, nil
}

func GetAllUnapprovedBlogsForSpace(spaceName string) ([]models.Blog, error) {
	//get unapproved blogs
	//return all blogs associated with spaceid that are unapproved
	//return all blogs associated with spaceid that are approved
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(spaceName),
			},
			":v2": {
				S: aws.String(models.False),
			},
		},
		KeyConditionExpression: aws.String("space_name = :v1 and approved = :v2"),
		TableName:              aws.String(models.BlogsTable),
		IndexName:              aws.String(models.GlobalIndexSpaceNameApproved),
	}

	result, err := db.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	blogs := []models.Blog{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &blogs)
	return blogs, nil
}

func GetBlogsByAuthorEmail(authorEmail string) ([]models.Blog, error) {
	//return all blogs by author email
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(authorEmail),
			},
		},
		KeyConditionExpression: aws.String("author_email = :v1"),
		TableName:              aws.String(models.BlogsTable),
		IndexName:              aws.String(models.GlobalIndexAuthorEmail),
	}

	result, err := db.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	blogs := []models.Blog{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &blogs)
	return blogs, nil
}
