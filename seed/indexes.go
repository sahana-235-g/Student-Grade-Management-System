package seed

import (
    "context"
    "log"
    "student-grade-api/config"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"           
    "go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Users: Email unique
    _, err := config.GetCollection("users").Indexes().CreateOne(ctx, mongo.IndexModel{
        Keys:    bson.D{{Key: "email", Value: 1}},
        Options: options.Index().SetUnique(true),
    })
    if err != nil {
        log.Println("Index error (users):", err)
    }

    // Enrollment: student & course index
    _, _ = config.GetCollection("enrollments").Indexes().CreateOne(ctx, mongo.IndexModel{
        Keys: bson.D{
            {Key: "student_id", Value: 1},
            {Key: "course_id", Value: 1},
        },
    })

    // Grade: student & course index
    _, _ = config.GetCollection("grades").Indexes().CreateOne(ctx, mongo.IndexModel{
        Keys: bson.D{
            {Key: "student_id", Value: 1},
            {Key: "course_id", Value: 1},
        },
    })

    log.Println("Indexes created successfully")
}