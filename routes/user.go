package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)



func UserRoutes(router *gin.Engine) {
    user := router.Group("/user")
    {
        user.GET("/", GetUsers)
        user.POST("/", CreateUser)
    }
}

var users = []map[string]string{
    {"name": "Codey"},
    {"name": "John"},
}

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "users": users,
    })
}

func CreateUser(c *gin.Context) {
    var user map[string]string
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    users = append(users, user) // add to slice

    c.JSON(http.StatusCreated, gin.H{
        "message": "User created",
        "user":    user,
    })
}
