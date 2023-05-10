package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"../models"
)

var db *gorm.DB

func Project(c *gin.Context) {
	c.HTML(http.StatusOK, "project.html", nil)
}

func ProjectPost(c *gin.Context) {
	// Check if user is logged in
	user, loggedIn := c.Get("user")
	if !loggedIn {
		// If not logged in, redirect to login page
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// Get form values
	projectType := c.PostForm("project_type")
	projectGenre := c.PostForm("project_genre")

	// Create new project
	project := models.Project{
		Type:  projectType,
		Genre: projectGenre,
		UserID: user.(uint),  // assuming "user" in context is the user ID
	}

	// Save the project to the database
	db.Create(&project)

	c.Redirect(http.StatusSeeOther, "/")
}
