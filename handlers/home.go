package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	// Check if user is logged in
	if _, loggedIn := c.Get("user"); !loggedIn {
		// If not logged in, redirect to login page
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	c.HTML(http.StatusOK, "home.html", nil)
}
