package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"../models"
)

var db *gorm.DB

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the session
		userID, exists := c.Get("user")
		if !exists {
			// If no user ID, redirect to login page
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		// Fetch the user from the database
		var user models.User
		if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
			// If user not found, redirect to login page
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		// User is authenticated, proceed with the request
		c.Next()
	}
}
