package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}

func LoadRouter(webConfig map[string]string) {
	// Create a new Gin router
	router := gin.Default()
	gin.ForceConsoleColor()
	router.Use(gin.Recovery())
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	// Endpoint to refresh the JWT token
	router.POST("/token/refresh", refreshHandler(webConfig))
	router.POST("/login", loginHandler(webConfig))

	// Endpoint to authenticate with the JWT token
	router.POST("/token/authenticate", authenticateHandler(webConfig), authMiddleware(webConfig))

	// Endpoint to check the server's health
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/api/host/all", Hosts(webConfig))
	router.GET("/api/host/get/:hostID", Host(webConfig))

	// Start the server
	router.Run(fmt.Sprint(":", webConfig["port"]))
}
