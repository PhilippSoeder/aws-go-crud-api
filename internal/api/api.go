package api

import (
	"encoding/json"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/philippsoeder/aws-go-crud-api/internal/db"
	"github.com/philippsoeder/aws-go-crud-api/internal/models"
	"github.com/philippsoeder/aws-go-crud-api/pkg/logger"
)

var log *slog.Logger

func init() {
	log = logger.GetLogger()
	log.Debug("Initializing API finished.")
}

func HandleGetAllNotes() (events.APIGatewayProxyResponse, error) {
	result, err := db.GetAllNotes()
	if err != nil {
		log.Error("Error getting notes", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal server error",
		}, nil
	}
	jsonBody, err := json.Marshal(result)
	if err != nil {
		log.Error("Error marshalling JSON", "error", err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonBody),
	}, nil
}

func HandleGetNoteByID(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	noteID := request.PathParameters["note-id"]
	result, err := db.GetNoteByID(noteID)
	if err != nil {
		log.Error("Error getting note", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal server error",
		}, nil
	}
	if result.ID == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Note not found",
		}, nil
	}
	jsonBody, err := json.Marshal(result)
	if err != nil {
		log.Error("Error marshalling JSON", "error", err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonBody),
	}, nil
}

func HandleInsertNote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var note models.Note
	err := json.Unmarshal([]byte(request.Body), &note)
	if err != nil {
		log.Error("Error unmarshalling JSON", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Bad request",
		}, nil
	}
	note, err = db.InsertNote(note)
	if err != nil {
		log.Error("Error inserting note", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal server error",
		}, nil
	}
	jsonbody, err := json.Marshal(note)
	if err != nil {
		log.Error("Error marshalling JSON", "error", err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       string(jsonbody),
	}, nil
}

func HandleUpdateNoteByID(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var note models.Note
	err := json.Unmarshal([]byte(request.Body), &note)
	if err != nil {
		log.Error("Error unmarshalling JSON", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Bad request",
		}, nil
	}
	noteID := request.PathParameters["note-id"]
	err = db.UpdateNoteByID(noteID, note)
	if err != nil {
		log.Error("Error updating note", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal server error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Success",
	}, nil
}

func HandleDeleteNoteByID(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	noteID := request.PathParameters["note-id"]
	err := db.DeleteNoteByID(noteID)
	// FIXME: currently there is no error if note is not existing
	if err != nil {
		log.Error("Error deleting note", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal server error",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Success",
	}, nil
}
