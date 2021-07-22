package app

import (
	"github.com/hidayatarg/go-microservice/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}