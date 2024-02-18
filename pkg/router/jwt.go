package router

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("")

func createToken(webConfig map[string]string, username string, roles []string) (string, error) {
	claims := &Claims{
		Username: username,
		Roles:    roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(webConfig["jwtKey"])
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func authenticateHandler(webConfig map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len("Bearer "):]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Authorization header not provided"})
			c.Abort()
			return
		}

		// Replace "your_secret_key" with your actual secret key
		secretKey := webConfig["jwtKey"]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Invalid token"})
			c.Abort()
			return
		}

		expirationTime := int64(claims["exp"].(float64))
		if time.Now().Unix() > expirationTime {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Token has expired"})
			c.Abort()
			return
		}

		// If token is valid, you can proceed with the request
		c.Next()
	}
}

// swagger:route POST /api/refresh refresh refreshHandler
//
// Refreshes the data in the system.
//
// This endpoint will refresh the data in the system based on the configuration settings
// provided in the environment variables.
//
// Responses:
//
//	200: successResponse
//	400: errorResponse
//	500: errorResponse
func refreshHandler(webConfig map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the username from the existing token
		claims := &Claims{}
		tokenString := c.GetHeader("Authorization")[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return webConfig["jwtKey"], nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Invalid or expired token"})
			return
		}
		// Create a new token with the same username
		newTokenString, err := createToken(webConfig, claims.Username, claims.Roles)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Error creating new token"})
			return
		}
		// Return the new token to the client
		c.JSON(http.StatusOK, gin.H{"token": newTokenString})
	}
}

// Middleware function to verify the JWT token
func authMiddleware(webConfig map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Authorization header required"})
			return
		}
		tokenString := authHeader[len("Bearer "):]

		// Parse the token using the secret key
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return webConfig["jwtKey"], nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Invalid or expired token"})
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
