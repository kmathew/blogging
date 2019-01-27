package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"blogging/models"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-west-1").WithEndpoint("http://127.0.0.1:8000"))
//var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1").WithEndpoint("http://127.0.0.1:8000"))

func registerAuthor(name string, displayName string, email string) (string, error) {
	author := models.Author {
		Name: name,
		DisplayName: displayName,
		Email: email,
	}

	// marshal the movie struct into an aws attribute value
	authorAVMap, err := dynamodbattribute.MarshalMap(author)
	if err != nil {
		panic("Cannot marshal author into AttributeValue map")
	}

	// create the api params
	params := &dynamodb.PutItemInput{
		TableName: aws.String("Authors"),
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

func createSpace(name string, ownerEmail string) error {

	space := models.Space {
		Name: name,
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

func getAuthor(email string, displayName string) (models.Author, error) {
	var author models.Author
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Authors"),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
			"display_name": {
				S: aws.String(displayName),
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

func getAuthorSpace(authorEmail string) ([]models.Space, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(authorEmail),
			},
		},
		KeyConditionExpression: aws.String("owner_email = :v1"),
		ProjectionExpression:   aws.String(models.SpaceName),
		TableName:              aws.String(models.SpacesTable),
		IndexName: aws.String(models.GlobalIndexOwnerEmail),
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

func createBlog(title string, content []byte, spaceName string, authorEmail string) {
	//generate id
	//insert blog into table
	//if spaceid belongs to authorid.... then auto approve
	//if not then send blog for approval

	blog := models.Blog{
		Title: title,
		Content: content,
		SpaceName: spaceName,
		AuthorEmail: authorEmail,
		Approved: models.False,
	}

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
		return
	}

	// print the response data
	fmt.Println(resp)
}

func getBlog(title string) ([]models.Blog, error) {
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

func approveBlog(spaceName string, blogTitle string, approverEmail string) {
	blogsList, err := getBlog(blogTitle)
	if err != nil {

	}
	spacesList, err := getAuthorSpace(approverEmail)

	if (blogsList[0].SpaceName == spaceName && spacesList[0].Name == spaceName) {
		//means we can approve it
	} else {
		//fail.. cant approve it because blog/space doesnt exist or not owner of space
	}

}

func getBlogsForSpaceName(spaceName string) ([]models.Blog, error) {
	//return all blogs associated with spaceid that are approved
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(spaceName),
			},
		},
		KeyConditionExpression: aws.String("space_name = :v1"),
		TableName:              aws.String(models.BlogsTable),
		IndexName: aws.String(models.GlobalIndexSpaceName),
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

func getApprovedBlogs() ([]models.Blog, error) {
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
		IndexName: aws.String(models.LocalIndexApproved),
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

func getUnapprovedBlogs() ([]models.Blog, error) {
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
		IndexName: aws.String(models.GlobalIndexApproved),
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

func getBlogsByAuthorEmail(authorEmail string) ([]models.Blog, error) {
	//return all blogs by author email
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(authorEmail),
			},
		},
		KeyConditionExpression: aws.String("author_email = :v1"),
		TableName:              aws.String(models.BlogsTable),
		IndexName: aws.String(models.GlobalIndexAuthorEmail),
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

func main() {
	/*err := registerAuthor("kevin", "tile", "yo@yo.com")
	if(err != nil ) {
		fmt.Println("Failed to register author")
	}*/
	a, err := getAuthor("yo@yo.com", "tile")
	if (err != nil) {
		panic("RIP")
	}

	//createSpace("FUNZONE", "yo@yo.com")

	fmt.Println(a.Email)

	s, err := getAuthorSpace("yo@yo.com")
	if (err != nil) {
		panic("RIP getAuthorSpace")
	}
	fmt.Println(s[0].Name)
	createBlog("fun2",[]byte("mkmkk"),"FUNZONE","yo@yo.com")

	/*b, err := getBlog("fun1")
	if (err != nil) {
		panic("RIP make blog")
	}*/

	b, err := getUnapprovedBlogs()
	if (err != nil) {
		panic("RIP unapproved blog")
	}

	fmt.Println(b)
}