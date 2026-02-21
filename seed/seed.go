package seed

import (
    "context"
    "fmt"
    "student-grade-api/config"
    "student-grade-api/models"
    "student-grade-api/utils"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func RunSeed() {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    userCol := config.GetCollection("users")
    courseCol := config.GetCollection("courses")
    enrollCol := config.GetCollection("enrollments")
    gradeCol := config.GetCollection("grades")

    // Clear existing data (optional)
    userCol.DeleteMany(ctx, bson.M{})
    courseCol.DeleteMany(ctx, bson.M{})
    enrollCol.DeleteMany(ctx, bson.M{})
    gradeCol.DeleteMany(ctx, bson.M{})

    // === Users ===
    adminPass, _ := utils.HashPassword("admin123")
    teacherPass, _ := utils.HashPassword("teacher123")
    studentPass, _ := utils.HashPassword("student123")

    admin := models.User{
        ID:       primitive.NewObjectID(),
        Name:     "Admin User",
        Email:    "admin@test.com",
        Password: adminPass,
        Role:     "admin",
    }

    teacher := models.User{
        ID:       primitive.NewObjectID(),
        Name:     "Teacher User",
        Email:    "teacher@test.com",
        Password: teacherPass,
        Role:     "teacher",
    }

    student := models.User{
        ID:       primitive.NewObjectID(),
        Name:     "Student User",
        Email:    "student@test.com",
        Password: studentPass,
        Role:     "student",
    }

    userCol.InsertOne(ctx, admin)
    userCol.InsertOne(ctx, teacher)
    userCol.InsertOne(ctx, student)

    // === Courses ===
    course1 := models.Course{
        ID:        primitive.NewObjectID(),
        Name:      "Mathematics",
        Code:      "MATH101",
        Credits:   4,
        TeacherID: teacher.ID,
    }

    course2 := models.Course{
        ID:        primitive.NewObjectID(),
        Name:      "Physics",
        Code:      "PHYS101",
        Credits:   3,
        TeacherID: teacher.ID,
    }

    courseCol.InsertOne(ctx, course1)
    courseCol.InsertOne(ctx, course2)

    // === Enrollments (Student enrolled in both courses) ===
    enrollment1 := models.Enrollment{
        ID:        primitive.NewObjectID(),
        StudentID: student.ID,
        CourseID:  course1.ID,
    }

    enrollment2 := models.Enrollment{
        ID:        primitive.NewObjectID(),
        StudentID: student.ID,
        CourseID:  course2.ID,
    }

    enrollCol.InsertOne(ctx, enrollment1)
    enrollCol.InsertOne(ctx, enrollment2)

    // === Grades ===
    grade1 := models.Grade{
        ID:        primitive.NewObjectID(),
        StudentID: student.ID,
        CourseID:  course1.ID,
        Grade:     85,
    }

    grade2 := models.Grade{
        ID:        primitive.NewObjectID(),
        StudentID: student.ID,
        CourseID:  course2.ID,
        Grade:     90,
    }

    gradeCol.InsertOne(ctx, grade1)
    gradeCol.InsertOne(ctx, grade2)

    fmt.Println("🌱 Database seeded successfully!")
}