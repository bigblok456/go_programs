// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.LoadHTMLFiles("templates\\index.html")
// 	r.GET("/vullipuka", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", gin.H{})
// 	})
// 	r.Run("localhost:8080")
// }
