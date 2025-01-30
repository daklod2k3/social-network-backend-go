package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadRepository represents a read-only repository interface
type ReadRepository[T any] struct {
	collection *mongo.Collection
}

// NewReadRepository creates a new ReadRepository
func NewReadRepository[T any](collection *mongo.Collection) *ReadRepository[T] {
	return &ReadRepository[T]{
		collection: collection,
	}
}

// FindByID retrieves a document by its ID
func (r *ReadRepository[T]) FindByID(ctx context.Context, id interface{}) (*T, error) {
	var entity T
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll retrieves all documents matching a filter
func (r *ReadRepository[T]) FindAll(ctx context.Context, filter bson.D, opts ...*options.FindOptions) ([]T, error) {
	var entities []T
	cursor, err := r.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var entity T
		if err := cursor.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
