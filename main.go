package main

import (
	"fmt"
	services "go-task/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Configurations struct {
	Server          ServerConfigurations
	REQUEST_TIMEOUT int
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
	env  string
}

func getJoke(c *gin.Context) {
	response := services.GetOrientedJoke()
	if response != "" {
		c.String(http.StatusOK, response)
	} else {
		c.String(http.StatusInternalServerError, "Unable to retrive joke")
	}
}

func main() {
	fmt.Println("API initialization start...")
	// set configuration from config file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration Configurations
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)

	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	if viper.GetString("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/joke", getJoke)
	fmt.Println("API ready...")
	router.Run("localhost:" + viper.GetString("server.port"))
}
