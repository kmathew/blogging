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
	case "GET":
		q := req.QueryStringParameters
		if q[models.OwnerEmail] != "" {
			return showByOwner(req)
		} else if q[models.SpaceName] != "" {
			return showBySpaceName(req)
		} else {
			return clientError(http.StatusBadRequest)
		}
	case "POST":
		return create(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func showBySpaceName(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	spaceName := req.QueryStringParameters[models.SpaceName]

	space, err := db.GetSpaceByName(spaceName)
	if err != nil {
		return serverError(err)
	}
	if space == nil {
		return clientError(http.StatusNotFound)
	}

	js, err := json.Marshal(space)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func showByOwner(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ownerEmail := req.QueryStringParameters[models.OwnerEmail]

	space, err := db.GetAuthorSpace(ownerEmail)
	if err != nil {
		return serverError(err)
	}
	if space == nil {
		return clientError(http.StatusNotFound)
	}

	js, err := json.Marshal(space)
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

	space := new(models.Space)
	err := json.Unmarshal([]byte(req.Body), space)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if space.Name == "" || space.OwnerEmail == "" {
		return clientError(http.StatusBadRequest)
	}

	err = db.CreateSpace(space.Name, space.OwnerEmail)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/space?space_name=%s", space.Name)},
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
