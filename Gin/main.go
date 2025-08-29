package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default() // Sets up a router with default middleware
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    r.Run() // Runs on localhost:8080
}
