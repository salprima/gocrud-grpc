package repository

import (
	"context"
	"log"
	"time"

	"github.com/salprima/gocrud-grpc/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User repository
type UserRepo struct {
	db  *mongo.Database
	col *mongo.Collection
}

// Instantiate new UserRepo
func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		db:  db,
		col: db.Collection(model.UserCollection),
	}
}

// Save single user
func (r *UserRepo) Save(u *model.User) (model.User, error) {
	log.Printf("Save(%v) \n", u)
	ctx, cancel := timeoutContext()
	defer cancel()

	var user model.User
	res, err := r.col.InsertOne(ctx, u)
	if err != nil {
		log.Println(err)
		return user, err
	}

	err = r.col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&user)
	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

// Find user by its id
func (r *UserRepo) FindByID(id string) (model.User, error) {
	log.Printf("FindByID(%s) \n", id)
	ctx, cancel := timeoutContext()
	defer cancel()

	var user model.User
	oid, _ := primitive.ObjectIDFromHex(id)
	err := r.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

// Find all user
func (r *UserRepo) FindAll() ([]model.User, error) {
	log.Println("FindAll()")
	ctx, cancel := timeoutContext()
	defer cancel()

	var users []model.User
	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return users, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

// Delete user by its id
func (r *UserRepo) DeleteByID(id string) (bool, error) {
	log.Printf("DeleteByID(%s) \n", id)
	ctx, cancel := timeoutContext()
	defer cancel()

	var user model.User
	oid, _ := primitive.ObjectIDFromHex(id)
	err := r.col.FindOneAndDelete(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		log.Printf("Fail to delete user: %v \n", err)
		return false, err
	}
	log.Printf("Deleted_User(%v) \n", user)
	return true, nil
}

// creating context background with timeout 60 seconds
func timeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
}
