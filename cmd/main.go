package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/philippsoeder/aws-go-crud-api/internal/api"
	"github.com/philippsoeder/aws-go-crud-api/pkg/logger"
)

var log *slog.Logger

func init() {
	log = logger.GetLogger()
	log.Debug("Initializing Lambda function finished.")
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Debug("Request", "request", request, "context", ctx)

	switch request.HTTPMethod {
	case "GET":
		if request.Resource == "/notes" {
			return api.HandleGetAllNotes()
		} else if request.Resource == "/notes/{note-id}" {
			return api.HandleGetNoteByID(request)
		}
	case "POST":
		return api.HandleInsertNote(request)
	case "PUT":
		return api.HandleUpdateNoteByID(request)
	case "DELETE":
		return api.HandleDeleteNoteByID(request)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 405,
		Body:       "Method not allowed",
	}, nil
}

func main() {
	lambda.Start(handler)
}
