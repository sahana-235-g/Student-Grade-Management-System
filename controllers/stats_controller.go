package controllers

import (
    "context"
    "net/http"
    "student-grade-api/config"
    "student-grade-api/models"
    "student-grade-api/services"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// ---------------- GPA ---------------------
func GetGPA(c *gin.Context) {
    gradeCol := config.GetCollection("grades") // FIXED

    studentId := c.Param("studentId")
    oid, err := primitive.ObjectIDFromHex(studentId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := gradeCol.Find(ctx, bson.M{"student_id": oid})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch grades"})
        return
    }

    var grades []models.Grade
    if err := cursor.All(ctx, &grades); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading grades"})
        return
    }

    gpa := services.CalculateGPA(grades)

    c.JSON(http.StatusOK, gin.H{"gpa": gpa})
}

// ------------- COURSE AVERAGE -------------
func GetCourseAverage(c *gin.Context) {
    gradeCol := config.GetCollection("grades") // FIXED

    courseId := c.Param("courseId")
    oid, err := primitive.ObjectIDFromHex(courseId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := gradeCol.Find(ctx, bson.M{"course_id": oid})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch grades"})
        return
    }

    var grades []models.Grade
    if err := cursor.All(ctx, &grades); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading grades"})
        return
    }

    avg := services.CalculateCourseAverage(grades)

    c.JSON(http.StatusOK, gin.H{"average": avg})
}