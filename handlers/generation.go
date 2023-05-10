package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
	"./handlers"
	"./middleware"
	"./models"
)

var db *gorm.DB

func init() {
	var err error
	// Connect to your database here
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})
}

func main() {
	router := gin.Default()

	// Middleware
	router.Use(middleware.Session())

	// Routes
	router.GET("/", handlers.Home)
	router.GET("/login", handlers.Login)
	router.POST("/login", handlers.LoginPost)
	router.GET("/logout", handlers.Logout)
	router.GET("/project", handlers.Project)
	router.POST("/project", handlers.ProjectPost)
	router.GET("/generation", handlers.Generation)
	router.POST("/generation", handlers.GenerationPost)

	// Start server
	router.Run(":8080")
}
