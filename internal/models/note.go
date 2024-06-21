package models

type Note struct {
	ID        string `json:"id" dynamo:"ID,hash"`
	Title     string `json:"title" dynamo:"title"`
	Content   string `json:"content" dynamo:"content"`
	CreatedAt string `json:"createdAt" dynamo:"created_at"`
	UpdatedAt string `json:"updatedAt" dynamo:"updated_at"`
}
