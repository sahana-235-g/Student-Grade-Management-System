package controllers

import (
    "context"
    "net/http"
    "strconv"
    "student-grade-api/config"
    "student-grade-api/models"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// AssignGradeRequest - request body for assigning grade
type AssignGradeRequest struct {
    StudentID string      `json:"student_id" binding:"required"`
    CourseID  string      `json:"course_id" binding:"required"`
    Grade     interface{} `json:"grade" binding:"required"` // accepts number or string
}

// Teacher assigns grade
func AssignGrade(c *gin.Context) {
    gradeCol := config.GetCollection("grades")

    var req AssignGradeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    studentOID, err := primitive.ObjectIDFromHex(req.StudentID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
        return
    }

    courseOID, err := primitive.ObjectIDFromHex(req.CourseID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    var gradeVal float64
    switch v := req.Grade.(type) {
    case float64:
        gradeVal = v
    case string:
        gradeVal, err = strconv.ParseFloat(v, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grade value"})
            return
        }
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grade value"})
        return
    }

    grade := models.Grade{
        ID:        primitive.NewObjectID(),
        StudentID: studentOID,
        CourseID:  courseOID,
        Grade:     gradeVal,
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = gradeCol.InsertOne(ctx, grade)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Could not assign grade"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Grade assigned"})
}

// Teacher updates grade
func UpdateGrade(c *gin.Context) {
    gradeCol := config.GetCollection("grades") // FIXED

    var grade models.Grade

    if err := c.ShouldBindJSON(&grade); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{
        "student_id": grade.StudentID,
        "course_id":  grade.CourseID,
    }

    update := bson.M{
        "$set": bson.M{
            "grade": grade.Grade,
        },
    }

    _, err := gradeCol.UpdateOne(ctx, filter, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update grade"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Grade updated"})
}

// GradeResponse - API response with string IDs for frontend
type GradeResponse struct {
    ID        string  `json:"id"`
    StudentID string  `json:"student_id"`
    CourseID  string  `json:"course_id"`
    Grade     float64 `json:"grade"`
}

// Get student's grades
func GetGrades(c *gin.Context) {
    gradeCol := config.GetCollection("grades")

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

    resp := make([]GradeResponse, len(grades))
    for i, g := range grades {
        resp[i] = GradeResponse{
            ID:        g.ID.Hex(),
            StudentID: g.StudentID.Hex(),
            CourseID:  g.CourseID.Hex(),
            Grade:     g.Grade,
        }
    }
    c.JSON(http.StatusOK, resp)
}