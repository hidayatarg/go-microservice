package services

import (
	"github.com/hidayatarg/go-microservice/mvc/domain"
)

func GetUser(userId int64) (domain.User, error){
	return domain.GetUser(userId)
	
}