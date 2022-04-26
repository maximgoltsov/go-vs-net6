package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.GET("/", getHome)

	router.Run(":8888")
}

func getHome(c *gin.Context) {
	time.Sleep(100 * time.Millisecond)
	c.JSON(200, "Ok")
}
