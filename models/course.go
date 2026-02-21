package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name      string             `bson:"name" json:"name"`
    Code      string             `bson:"code" json:"code"`
    Credits   int                `bson:"credits" json:"credits"`
    TeacherID primitive.ObjectID `bson:"teacher_id,omitempty" json:"teacher_id,omitempty"`
}