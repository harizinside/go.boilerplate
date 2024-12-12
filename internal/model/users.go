package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdateAt  time.Time          `bson:"update_at,omitempty"`
}
