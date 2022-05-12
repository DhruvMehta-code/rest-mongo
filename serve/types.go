package serve

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Desription string             `json:"desc,omitempty" bson:"desc,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at,omitempty"`
}
