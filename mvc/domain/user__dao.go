package domain

import (
	"fmt"
	"github.com/hidayatarg/go-microservice/mvc/utils"
	"net/http"
)

// datastore
var (
	users = map[int64]*User {
		123: { Id: 123, FirstName: "Hidayet", LastName: "Arghandabi", Email: "myEmail@abc.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {

	if user := users[userId]; user != nil {
		return user, nil

	}
	return nil, &utils.ApplicationError {
		Message: fmt.Sprintf("User %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not found",
	}
}