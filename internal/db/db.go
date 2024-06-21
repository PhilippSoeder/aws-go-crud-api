package db

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/philippsoeder/aws-go-crud-api/internal/models"
	"github.com/philippsoeder/aws-go-crud-api/pkg/logger"
)

var log *slog.Logger
var ddbClient *dynamodb.Client
var notesTable string

func init() {
	log = logger.GetLogger()
	ddbRegion := os.Getenv("DYNAMODB_REGION")
	notesTable = os.Getenv("DYNAMODB_TABLE_NAME")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(ddbRegion))
	if err != nil {
		log.Error("unable to load SDK config.", "error", err)
	}
	ddbClient = dynamodb.NewFromConfig(cfg)
	log.Debug("Initializing DynamoDB table finished.")
}

func GetAllNotes() ([]models.Note, error) {
	var notes []models.Note
	scan, err := ddbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: &notesTable,
	})
	if err != nil {
		log.Error("Got error calling Scan", "error", err)
	}
	err = attributevalue.UnmarshalListOfMaps(scan.Items, &notes)
	return notes, err
}

func GetNoteByID(id string) (*models.Note, error) {
	var note models.Note
	result, err := ddbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &notesTable,
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		log.Error("Got error calling GetItem", "error", err)
		return nil, err
	}
	err = attributevalue.UnmarshalMap(result.Item, &note)
	return &note, err
}

func InsertNote(note models.Note) (models.Note, error) {
	if note.ID == "" {
		note.ID = uuid.New().String()
	}
	note.CreatedAt = time.Now().String()
	note.UpdatedAt = note.CreatedAt
	item, err := attributevalue.MarshalMap(note)
	if err != nil {
		return note, err
	}
	_, err = ddbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &notesTable,
		Item:      item,
	})
	return note, err
}

func UpdateNoteByID(id string, note models.Note) error {
	updateExpression := "SET "
	expressionAttributeValues := make(map[string]types.AttributeValue)
	if note.Title != "" {
		updateExpression += "Title = :t, "
		expressionAttributeValues[":t"] = &types.AttributeValueMemberS{Value: note.Title}
	}
	if note.Content != "" {
		updateExpression += "Content = :c, "
		expressionAttributeValues[":c"] = &types.AttributeValueMemberS{Value: note.Content}
	}
	updateExpression += "UpdatedAt = :u"
	expressionAttributeValues[":u"] = &types.AttributeValueMemberS{Value: time.Now().String()}
	_, err := ddbClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &notesTable,
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression:          &updateExpression,
		ExpressionAttributeValues: expressionAttributeValues,
	})
	return err
}

func DeleteNoteByID(id string) error {
	_, err := ddbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &notesTable,
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}
