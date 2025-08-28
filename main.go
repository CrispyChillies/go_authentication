package main

import (
	"net/http"

	"github.com/CrispyChillies/go_authentication/db"
	"github.com/CrispyChillies/go_authentication/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Go Authentication API"})
	})

	// Auth routes
	router.POST("/register", handlers.RegisterHandler(db.DB))
	router.POST("/login", handlers.LoginHandler(db.DB))

	router.Run(":8080")
}
