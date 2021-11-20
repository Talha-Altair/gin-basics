package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/altair", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"talha": "altair",
		})
	})

	main2()

	r.Run()
}
