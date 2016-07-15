package main

import "github.com/gin-gonic/gin"

import "github.com/gin-gonic/contrib/static"

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./html", true)))

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	r.Run() // listen and server on 0.0.0.0:8080
}
