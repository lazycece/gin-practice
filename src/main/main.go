package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/hello/:name", func(context *gin.Context) {
		name := context.Param("name")
		fmt.Printf("params name  = " + name)
		context.String(http.StatusOK, "hello, "+name)
	})
	router.Run(":8080")
}
