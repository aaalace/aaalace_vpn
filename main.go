package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func main() {
	router := gin.Default()
	router.GET("/ping", pong)

	err := router.Run(":1234")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
