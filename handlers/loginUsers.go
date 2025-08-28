package handlers

import (
	"database/sql"
	"net/http"

	"github.com/CrispyChillies/go_authentication/utils"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check user from DB
		var hashedPassword string
		var userId int
		err := db.QueryRow("SELECT id, password FROM users WHERE username=$1", req.Username).Scan(&userId, &hashedPassword)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}

		// Verify password
		if !utils.CheckPassword(req.Password, hashedPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}

		// Generate JWT
		token, err := utils.GenerateJWT(userId, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
