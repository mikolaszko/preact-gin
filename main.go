package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./assets/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "This is an about page...",
		})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
