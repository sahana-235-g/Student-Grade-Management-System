package controllers

import (
    "context"
    "net/http"
    "student-grade-api/config"
    "student-grade-api/models"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// ADMIN / TEACHER Creates Course
func CreateCourse(c *gin.Context) {
    courseCol := config.GetCollection("courses") // FIXED

    var course models.Course

    if err := c.ShouldBindJSON(&course); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    course.ID = primitive.NewObjectID()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := courseCol.InsertOne(ctx, course)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create course"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
}

// EnrollStudentRequest - request body for enrollment
type EnrollStudentRequest struct {
    StudentID string `json:"student_id" binding:"required"`
}

// Enroll Student to Course
func EnrollStudent(c *gin.Context) {
    enrollCol := config.GetCollection("enrollments")
    courseIDStr := c.Param("id")

    courseOID, err := primitive.ObjectIDFromHex(courseIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    var req EnrollStudentRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
        return
    }

    studentOID, err := primitive.ObjectIDFromHex(req.StudentID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
        return
    }

    enr := models.Enrollment{
        ID:        primitive.NewObjectID(),
        StudentID: studentOID,
        CourseID:  courseOID,
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = enrollCol.InsertOne(ctx, enr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Enrollment failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Student enrolled"})
}

// Get List of All Courses
func GetAllCourses(c *gin.Context) {
    courseCol := config.GetCollection("courses") // FIXED

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := courseCol.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch courses"})
        return
    }

    var courses []models.Course
    if err := cursor.All(ctx, &courses); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading courses"})
        return
    }

    c.JSON(http.StatusOK, courses)
}