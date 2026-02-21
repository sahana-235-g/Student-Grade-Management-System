package main

import (
    "log"
    "student-grade-api/config"
    "student-grade-api/controllers"
    "student-grade-api/middleware"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectMongo()
    // seed.RunSeed()

    router := gin.Default()

    // CORS - allow frontend to connect from different port
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "http://127.0.0.1:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // Auth Routes
    auth := router.Group("/auth")
    {
        auth.POST("/register", controllers.RegisterUser)
        auth.POST("/login", controllers.LoginUser)
    }

    api := router.Group("/api")
    api.Use(middleware.JWTAuthMiddleware())

    // RBAC routes
    userRoutes := api.Group("/users")
    {
        userRoutes.POST("/", middleware.RequireRole("admin"), controllers.CreateUser)
        userRoutes.GET("/", middleware.RequireRole("admin"), controllers.GetAllUsers)
    }

    courseRoutes := api.Group("/courses")
    {
        courseRoutes.POST("/", middleware.RequireRole("admin", "teacher"), controllers.CreateCourse)
        courseRoutes.POST("/:id/enroll", middleware.RequireRole("admin", "teacher"), controllers.EnrollStudent)
        courseRoutes.GET("/", controllers.GetAllCourses)
    }

    gradeRoutes := api.Group("/grades")
    {
        gradeRoutes.POST("/", middleware.RequireRole("teacher"), controllers.AssignGrade)
        gradeRoutes.PUT("/", middleware.RequireRole("teacher"), controllers.UpdateGrade)
        gradeRoutes.GET("/:studentId", middleware.RequireRole("admin", "teacher", "student"), controllers.GetGrades)
    }

    statsRoutes := api.Group("/stats")
    {
        statsRoutes.GET("/gpa/:studentId", controllers.GetGPA)
        statsRoutes.GET("/course-avg/:courseId", controllers.GetCourseAverage)
    }

    log.Println("Server running on port 8080")
    router.Run(":8080")
}