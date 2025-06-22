package models

import "time"

type Task struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Status      string    `json:"status" bson:"status"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
}
