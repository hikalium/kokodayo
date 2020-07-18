package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const host = "localhost"
const port = 8080
const staticPath = "./static"

func main() {
	router := gin.Default()
	router.GET("/api/hello/:id", func(c *gin.Context) {
		name := c.Param("id")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("kokodayo %s!", name),
		})
	})
	router.Run(fmt.Sprintf("%v:%v", host, port))
}
