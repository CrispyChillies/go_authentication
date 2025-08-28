package main

import (
	"net/http"

	"github.com/CrispyChillies/go_authentication/db"
	"github.com/CrispyChillies/go_authentication/handlers"
	"github.com/CrispyChillies/go_authentication/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	router := gin.Default()

	// Public routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Go Authentication API"})
	})

	router.POST("/register", handlers.RegisterHandler(db.DB))
	router.POST("/login", handlers.LoginHandler(db.DB))

	// Protected routes
	protected := router.Group("/api")
	protected.Use(utils.AuthMiddleware()) // <--- Áp middleware vào group này

	protected.GET("/profile", func(c *gin.Context) {
		userId, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{
			"message": "This is a protected route",
			"user_id": userId,
		})
	})

	router.Run(":8080")
}
