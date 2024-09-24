package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// for Json validation thats used for transmitting the data between the web applications
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func createBook(c *gin.Context) {
	var newbook book
	err := c.BindJSON(&newbook)
	if err != nil {
		return
	}
	books = append(books, newbook)
	c.IndentedJSON(http.StatusCreated, newbook)

}
func getBookbyID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
func bookbyID(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookbyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)

}
func getIndexofBOOK(id string) (int, error) {
	for i, t := range books {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("book not found")

}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	index, err := getIndexofBOOK(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "book not found"})
		return
	}

	books = append(books[:index], books[index+1:]...)
	c.IndentedJSON(http.StatusOK, gin.H{"msg": "book  is deleted"})

}
func checkoutBOOK(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "invalid id"})
		return
	}
	book, err := getBookbyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "book not found"})
		return

	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "book is not available"})
		return

	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookbyID)
	router.POST("/createbooks", createBook)
	router.DELETE("/delbooks/:id", deleteBook)
	router.PATCH("/checkout", checkoutBOOK)
	router.Run("localhost:8080")
}
