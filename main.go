package main

import (
	services "go-task/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getJoke(c *gin.Context) {
	response := services.GetOrientedJoke()
	if response != "" {
		c.String(http.StatusOK, response)
	} else {
		c.String(http.StatusInternalServerError, "Unable to retrive joke")
	}
}

func main() {
	router := gin.Default()
	router.GET("/joke", getJoke)
	router.Run("localhost:4001")
}
