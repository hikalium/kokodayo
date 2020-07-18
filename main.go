package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"main/kokodayo"
)

const host = "localhost"
const port = 8080
const staticPath = "./static"

type ItemRequest struct {
	ImageUrl string `json:"image_url" binding:"required"`
}

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("static", false)))
	router.GET("/api/hello/:id", func(c *gin.Context) {
		name := c.Param("id")
		c.JSON(200, kokodayo.GenHello(name))
	})
	router.POST("/api/items", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		var rq ItemRequest
		err := c.BindJSON(&rq)
		if err != nil {
			c.JSON(400, gin.H{"status": "failed to bind the request to json"})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "data": rq.ImageUrl})
	})
	router.Run(fmt.Sprintf("%v:%v", host, port))
}
