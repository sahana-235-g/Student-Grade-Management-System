package controllers

import (
    "context"
    "net/http"
    "student-grade-api/config"
    "student-grade-api/models"
    "student-grade-api/utils"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

// ADMIN CREATES USER
func CreateUser(c *gin.Context) {
    userCol := config.GetCollection("users") // FIXED

    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashed, _ := utils.HashPassword(user.Password)
    user.Password = hashed

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := userCol.InsertOne(ctx, user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// ADMIN GET ALL USERS
func GetAllUsers(c *gin.Context) {
    userCol := config.GetCollection("users") // FIXED

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := userCol.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
        return
    }

    var users []models.User

    if err := cursor.All(ctx, &users); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading users"})
        return
    }

    c.JSON(http.StatusOK, users)
}