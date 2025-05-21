package main

import (
    "github.com/gin-gonic/gin"
    "gin-rest-api/routes"
	"net/http"
)

func main() {
    router := gin.Default()

	router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
    })

    // Register routes
    routes.UserRoutes(router)
	

    router.Run(":8080") // server runs on localhost:8080
}
