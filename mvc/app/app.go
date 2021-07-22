package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	//router = gin.New()
	// the upper on doesnt use middleware and loggers
	router = gin.Default()
}

func StartApp() {

	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

