package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Enrollment struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    StudentID primitive.ObjectID `bson:"student_id" json:"student_id"`
    CourseID  primitive.ObjectID `bson:"course_id" json:"course_id"`
}