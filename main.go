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

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("static", false)))
	router.GET("/api/hello/:id", func(c *gin.Context) {
		name := c.Param("id")
		c.JSON(200, kokodayo.GenHello(name))
	})
	router.Run(fmt.Sprintf("%v:%v", host, port))
}
