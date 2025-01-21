package entity

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepo struct {
	collection *mongo.Collection
	read       *UserRepoRead
}

type UserRepoRead struct {
	collection *mongo.Collection
}

// init user repo database
func newRepo(schema *mongo.Database) *mongo.Collection {
	return schema.Collection("users")
}

// init repo can read and write
func NewRepo(schema *mongo.Database) *UserRepo {
	return &UserRepo{
		newRepo(schema),
		NewRepoRead(schema),
	}
}

// init repo only read
func NewRepoRead(schema *mongo.Database) *UserRepoRead {
	return &UserRepoRead{
		newRepo(schema),
	}
}

func (db *UserRepoRead) FindUser(id *uuid.UUID) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var user User
	err := db.collection.FindOne(ctx, bson.M{"user_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *UserRepo) CreateUser(id *uuid.UUID, name *string, avaPath *string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	doc := User{
		DisplayName: *name,
		UserId:      *id,
	}
	_, err := db.collection.InsertOne(ctx, &doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

//
//func (db *UserRepo) FindOrCreateUser(id *uuid.UUID, name *string) (*User, error) {
//	var user, err = db.read.FindUser(id)
//	switch {
//	case errors.Is(err, mongo.ErrNoDocuments):
//		return db.CreateUser(id, name, "")
//	case err != nil:
//		return nil, err
//	}
//	return user, nil
//}
