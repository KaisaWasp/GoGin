package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World")
	})

	fmt.Println("Привет")
	router.Run(":8080")
}
