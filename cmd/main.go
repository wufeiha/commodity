package main

import (
	"fmt"
	"go.uber.org/zap"
)

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorm/gorm"
	"github.com/spf13/viper"
)

func main() {
	// Initialize a new Gin router
	r := gin.Default()

	// Initialize Viper configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Initialize Zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize GORM
	_, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("Error initializing database, %s", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.Run()
}
