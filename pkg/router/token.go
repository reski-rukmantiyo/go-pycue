package router

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func loginHandler(webConfig map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Validate user credentials
		if !AuthenticateUser(user.Username, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// Generate JWT token
		token, err := GenerateToken(webConfig, user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Send token in response body
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func AuthenticateUser(username string, password string) bool {
	return true
}

func GenerateToken(webConfig map[string]string, username string) (string, error) {
	// Create a new token object with the claims for the user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(webConfig["jwtKey"])
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
