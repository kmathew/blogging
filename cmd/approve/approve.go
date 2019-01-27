package main

import (
	"encoding/json"
	"fmt"
	"github.com/kmathew/blogging/db"
	"github.com/kmathew/blogging/models"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "PUT":
		return create(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable)
	}

	approve := new(models.Approval)
	err := json.Unmarshal([]byte(req.Body), approve)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if approve.Title == "" || approve.SpaceName == "" || approve.ApproverEmail == "" {
		return clientError(http.StatusBadRequest)
	}

	str, err := db.ApproveBlog(approve.SpaceName, approve.Title, approve.ApproverEmail)
	if err != nil || str == "" {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 202,
		Headers:    map[string]string{"Location": fmt.Sprintf("/approve?title=%s", approve.Title)},
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
