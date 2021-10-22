package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Dob        string             `json:"dob" bson:"dob,omitempty"`
	Address    string             `json:"address" bson:"address,omitempty"`
	Desription string             `json:"description" bson:"description,omitempty"`
	CreatedAt  string             `json:"createdat" bson:"createdat,omitempty"`
}
