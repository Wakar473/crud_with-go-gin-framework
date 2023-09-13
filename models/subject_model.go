// models/subject_model.go

package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Subject struct {
    ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name   string             `json:"name"`
    Credits uint               `json:"credits"`
}
