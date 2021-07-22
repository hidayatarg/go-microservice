package services

import (
	"github.com/hidayatarg/go-microservice/mvc/domain"
	"github.com/hidayatarg/go-microservice/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError){
	// domain making query
	return domain.GetUser(userId)
}