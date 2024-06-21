package db

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/google/uuid"
	"github.com/guregu/dynamo/v2"
	"github.com/philippsoeder/aws-go-crud-api/internal/models"
	"github.com/philippsoeder/aws-go-crud-api/pkg/logger"
)

var notesTable dynamo.Table
var log *slog.Logger

func init() {
	log = logger.GetLogger()
	ddbRegion := os.Getenv("DYNAMODB_REGION")
	ddbTableName := os.Getenv("DYNAMODB_TABLE_NAME")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(ddbRegion))
	if err != nil {
		log.Error("unable to load SDK config.", "error", err)
	}
	ddb := dynamo.New(cfg)
	notesTable = ddb.Table(ddbTableName)
	log.Debug("Initializing DynamoDB table finished.")
}

func GetAllNotes() ([]models.Note, error) {
	var notes []models.Note
	err := notesTable.Scan().All(context.TODO(), &notes)
	return notes, err
}

func GetNoteByID(id string) (*models.Note, error) {
	var note models.Note
	err := notesTable.Get("ID", id).One(context.TODO(), &note)
	return &note, err
}

func InsertNote(note models.Note) (models.Note, error) {
	if note.ID == "" {
		note.ID = uuid.New().String()
	}
	note.CreatedAt = time.Now().String()
	note.UpdatedAt = note.CreatedAt
	err := notesTable.Put(note).Run(context.TODO())
	return note, err
}

func UpdateNoteByID(id string, note models.Note) error {
	upd := notesTable.Update("ID", id)
	newTitle := note.Title
	if newTitle != "" {
		upd.Set("title", newTitle)
	}
	newContent := note.Content
	if newContent != "" {
		upd.Set("content", newContent)
	}
	note.UpdatedAt = time.Now().String()
	upd.Set("updated_at", note.UpdatedAt)
	err := upd.Run(context.TODO())
	return err
}

func DeleteNoteByID(id string) error {
	err := notesTable.Delete("ID", id).Run(context.TODO())
	return err
}
