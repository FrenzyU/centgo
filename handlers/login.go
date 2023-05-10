package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"

	"../models"
)

var db *gorm.DB

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(c *gin.Context) {
	// Get form values
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Fetch the user from the database
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": "Invalid username or password"})
		return
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": "Invalid username or password"})
		return
	}

	// Log the user in by setting a value in the session
	c.Set("user", user.ID)

	c.Redirect(http.StatusSeeOther, "/")
}
