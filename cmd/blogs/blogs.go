package main

import (
	"encoding/json"
	"fmt"
	"github.com/kmathew/blogging/models"
	"github.com/kmathew/blogging/db"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		q := req.QueryStringParameters
		if q[models.Title] != "" {
			return showByTitle(req)
		} else if q[models.SpaceName] != "" {
			return showListBySpaceName(req)
		} else if q[models.AuthorEmail] != "" {
			return showListByAuthor(req)
		} else {
			return clientError(http.StatusBadRequest)
		}
	case "POST":
		return create(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func showByTitle(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	title := req.QueryStringParameters[models.Title]

	blog, err := db.GetBlog(title)
	if err != nil {
		return serverError(err)
	}
	if blog == nil {
		return clientError(http.StatusNotFound)
	}

	js, err := json.Marshal(blog)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func showListBySpaceName(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	spaceName := req.QueryStringParameters[models.SpaceName]

	blogs, err := db.GetBlogsForSpaceName(spaceName)
	if err != nil {
		return serverError(err)
	}
	if blogs == nil {
		return clientError(http.StatusNotFound)
	}

	js, err := json.Marshal(blogs)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func showListByAuthor(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authorEmail := req.QueryStringParameters[models.AuthorEmail]

	blogs, err := db.GetBlogsByAuthorEmail(authorEmail)
	if err != nil {
		return serverError(err)
	}
	if blogs == nil {
		return clientError(http.StatusNotFound)
	}

	js, err := json.Marshal(blogs)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable)
	}

	blog := new(models.Blog)
	err := json.Unmarshal([]byte(req.Body), blog)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if blog.Title == "" || blog.SpaceName == "" || blog.AuthorEmail == "" {
		return clientError(http.StatusBadRequest)
	}

	str, err := db.CreateBlog(blog.Title, blog.Content, blog.SpaceName, blog.AuthorEmail)
	if err != nil || str == "" {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/blog?title=%s", blog.Title)},
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(router)
}
