package entity

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"shared/database"
	"shared/entity"
	"time"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewRepo() *UserRepo {
	return &UserRepo{
		database.New().GetSchema().Collection("users"),
	}
}

func (db *UserRepo) FindUser(id *uuid.UUID) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var user entity.User
	err := db.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *UserRepo) CreateUser(id *uuid.UUID, name *string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	doc := entity.User{
		DisplayName: *name,
		UserId:      *id,
	}
	_, err := db.collection.InsertOne(ctx, &doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}
