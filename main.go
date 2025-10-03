package main

import (
	"go-currency/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/convert", handler.ConvertHandler)
	r.POST("/convert", handler.ConvertHandler)

	r.Run(":8080")
}
