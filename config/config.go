package config

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        uri = "mongodb://localhost:27017"
    }

    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("Mongo Client Error:", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal("Mongo Connect Error:", err)
    }

    DB = client.Database("student_grade_db")
    log.Println("Connected to MongoDB")
}

func GetCollection(name string) *mongo.Collection {
    return DB.Collection(name)
}