package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//
const UserCollection = "user"

//
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}
