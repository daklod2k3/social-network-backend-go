package database

import (
	"context"
	"fmt"
	"log"
	"shared/config"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	GetSchema() *mongo.Database
}

type service struct {
	db     *mongo.Client
	schema *mongo.Database
}

var (
	connectString = config.GetConfig().Database.ConnectString
)

func New() Service {
	fmt.Println("Connecting to", connectString)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectString))
	if err != nil {
		log.Fatal(err)
	}
	schema := client.Database("main")

	return &service{
		db:     client,
		schema: schema,
	}
}

func (s *service) GetSchema() *mongo.Database {
	return s.schema
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
