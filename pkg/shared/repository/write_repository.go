package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// WriteRepository represents a write-only repository interface
type WriteRepository[T any] struct {
	collection *mongo.Collection
}

// NewWriteRepository creates a new WriteRepository
func NewWriteRepository[T any](collection *mongo.Collection) *WriteRepository[T] {
	return &WriteRepository[T]{
		collection: collection,
	}
}

// Insert adds a new document to the collection
func (r *WriteRepository[T]) Insert(ctx context.Context, entity *T) (primitive.ObjectID, error) {
	//doc := bson.M{
	//	"$set": entity,
	//	"$currentDate": bson.M{
	//		"created_at": true,
	//	},
	//}
	//opts := options.Update().SetUpsert(true)
	//res, err := r.collection.UpdateOne(ctx, bson.M{
	//	"_id": primitive.NewObjectID(),
	//}, doc, opts)
	//if err != nil || res.UpsertedID == nil {
	//	return primitive.NilObjectID, err
	//}
	//return res.UpsertedID.(primitive.ObjectID), nil

	res, err := r.collection.InsertOne(ctx, entity)
	if err != nil || res.InsertedID == nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// Update updates a document by its ID
func (r *WriteRepository[T]) Update(ctx context.Context, id interface{}, update bson.M) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	return err
}

// Delete removes a document by its ID
func (r *WriteRepository[T]) Delete(ctx context.Context, id interface{}) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}
