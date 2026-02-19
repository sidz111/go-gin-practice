package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello world",
	})
}

type Student struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Contact string  `json:"contact"`
	Marks   float32 `json:"marks"`
}

type Book struct {
	Name          string `json:"name"`
	AuthorName    string `json:"author_name"`
	PublishedYear string `json:"published_year"`
	Price         int    `json:"price"`
}

func info(c *gin.Context) {
	var s Student
	c.BindJSON(&s)
	c.JSON(http.StatusOK, gin.H{
		"Name":    s.Name,
		"Address": s.Address,
		"Contact": s.Contact,
		"Marks":   s.Marks,
	})
}

func bookInfo(c *gin.Context) {
	var b Book
	c.BindJSON(&b)
	c.JSON(http.StatusOK, gin.H{
		"Name":          b.Name,
		"AuthorName":    b.AuthorName,
		"PublishedYear": b.PublishedYear,
		"Price":         b.Price,
	})
}

func main() {
	r := gin.Default()

	r.GET("/new/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	r.GET("/hello", hello)
	r.POST("/info", info)
	r.POST("/book", bookInfo)
	if err := r.Run(":9090"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
