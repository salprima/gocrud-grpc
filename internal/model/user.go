package model

import (
	"github.com/globalsign/mgo/bson"
)

//
const UserCollection = "user"

//
type User struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        `bson:"name"`
	Email string        `bson:"email"`
}
