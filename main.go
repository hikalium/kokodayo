package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"main/kokodayo"
	"os"
	"strconv"
)

const host = "localhost"
const port = 8080
const staticPath = "./static"

type Item struct {
	ItemId     int    `json:"item_id" db:"item_id"`
	TrainClass string `json:"image_url" db:"image_url"`
}

type ItemRequest struct {
	ImageUrl string `json:"image_url" binding:"required"`
}

var dbx *sqlx.DB

func main() {
	// Setup MYSQL connection
	var err error
	dbHost := os.Getenv("MYSQL_HOSTNAME")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	dbPort := os.Getenv("MYSQL_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	_, err = strconv.Atoi(dbPort)
	if err != nil {
		dbPort = "3306"
	}
	dbUser := os.Getenv("MYSQL_USER")
	if dbUser == "" {
		dbUser = "root"
	}
	dbname := os.Getenv("MYSQL_DATABASE")
	if dbname == "" {
		dbname = "kokodayo_test"
	}
	dbPass := os.Getenv("MYSQL_PASSWORD")
	if dbPass == "" {
		dbPass = "kokodayo"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbname,
	)

	dbx := sqlx.MustOpen("mysql", dsn)
	defer dbx.Close()
	log.Println("Connected db")

	// Setup routing
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
	router.GET("/api/items", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		itemList := []Item{}
		query := "SELECT * FROM items"
		err := dbx.Select(&itemList, query)
		if err != nil {
			log.Println("DB Error", err)
			c.JSON(500, gin.H{"status": "failed to query to DB"})
			return
		}

		c.JSON(200, gin.H{"status": "ok", "items": itemList})
	})
	router.Run(fmt.Sprintf("%v:%v", host, port))
}
