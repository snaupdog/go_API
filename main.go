package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "net/http"
)

type Book struct {
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {Title: "Book 1", Author: "nilesh 1"},
    {Title: "Book 2", Author: "Author 2"},
    {Title: "Book 3", Author: "Author 3"},
}

func main() {
    router := gin.Default()
    router.Use(cors.Default())

    // Endpoint to retrieve list of books
    router.GET("/books", func(c *gin.Context) {
        c.JSON(http.StatusOK, books)
    })

    router.Run(":8080")
}