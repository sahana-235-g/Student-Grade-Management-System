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
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// REGISTER USER
func RegisterUser(c *gin.Context) {
    userCollection := config.GetCollection("users") // FIXED

    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password
    hashedPassword, _ := utils.HashPassword(user.Password)
    user.Password = hashedPassword

    user.ID = primitive.NewObjectID()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := userCollection.InsertOne(ctx, user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LOGIN USER
func LoginUser(c *gin.Context) {
    userCollection := config.GetCollection("users") // FIXED

    var input models.User
    var dbUser models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := userCollection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&dbUser)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if !utils.CheckPasswordHash(input.Password, dbUser.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, _ := utils.GenerateJWT(dbUser.ID.Hex(), dbUser.Email, dbUser.Role)

    c.JSON(http.StatusOK, gin.H{"token": token})
}