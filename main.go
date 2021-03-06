package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"html/template"
	"log"
	"main/kokodayo"
	"os"
	"strconv"
)

const staticPath = "./static"

type Box struct {
	BoxId    int          `json:"box_id" db:"box_id"`
	ImageUrl template.URL `json:"image_url" db:"image_url"`
}

type Item struct {
	ItemId   int          `json:"item_id" db:"item_id"`
	BoxId    int          `json:"box_id" db:"box_id"`
	ImageUrl template.URL `json:"image_url" db:"image_url"`
}

type ItemRequest struct {
	BoxID    int    `json:"box_id" binding:"required"`
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
	log.Printf("dbPort = %v", dbPort)
	log.Printf("dbUser = %v", dbUser)

	dbx := sqlx.MustOpen("mysql", dsn)
	defer dbx.Close()
	log.Println("Connected db")

	// Setup routing
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
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
		log.Printf("Add an item to box id %v", rq.BoxID)
		dbx.MustExec("INSERT INTO items (box_id, image_url) VALUES (?, ?)", rq.BoxID, rq.ImageUrl)
		c.JSON(200, gin.H{"status": "ok"})
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
	router.GET("/api/boxes", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		boxList := []Box{}
		query := "SELECT * FROM boxes"
		err := dbx.Select(&boxList, query)
		if err != nil {
			log.Println("DB Error", err)
			c.JSON(500, gin.H{"status": "failed to query to DB"})
			return
		}

		c.JSON(200, gin.H{"status": "ok", "boxes": boxList})
	})
	router.GET("/b", func(c *gin.Context) {
		boxList := []Box{}
		query := "SELECT * FROM boxes"
		err := dbx.Select(&boxList, query)
		if err != nil {
			log.Println("DB Error", err)
			c.JSON(500, gin.H{"status": "failed to query to DB"})
			return
		}
		c.HTML(200, "box_list.html", gin.H{"boxes": boxList})
	})
	router.GET("/b/:id", func(c *gin.Context) {
		boxID := c.Param("id")
		var box Box
		err := dbx.Get(&box, "SELECT * FROM boxes WHERE box_id = ?", boxID)
		if err == sql.ErrNoRows {
			c.String(404, "Not found")
			return
		}
		var items []Item
		err = dbx.Select(&items, "SELECT * FROM items WHERE box_id = ?", boxID)
		c.HTML(200, "box_details.html", gin.H{"items": items, "box_id": boxID})
	})
	wwwPort := os.Getenv("WWW_PORT")
	if wwwPort == "" {
		wwwPort = "8080"
	}
	_, err = strconv.Atoi(wwwPort)
	if err != nil {
		wwwPort = "8080"
	}
	router.Run(fmt.Sprintf(":%v", wwwPort))
}
