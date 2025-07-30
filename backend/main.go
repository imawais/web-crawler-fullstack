package main

import (
    "github.com/gin-gonic/gin"
    "github.com/imawais/web-crawler-fullstack/backend/database"
	  "github.com/imawais/web-crawler-fullstack/backend/handlers"
)

func main() {
	if err := database.Connect(); err != nil {
		panic(err)
	}

	r := gin.Default()

	// Test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// API routes
	api := r.Group("/api")
	{
		api.POST("/urls", handlers.AddURL)
		api.GET("/urls", handlers.ListURLs)
	}

	r.Run(":8080")
}
