package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User {
		123: { Id: 123, FirstName: "Hidayet", LastName: "Arghandabi", Email: "myEmail@abc.com"},
	}
)

func GetUser(userId int64) (*User, error) {

	if user := users[userId]; user != nil {
		return user, nil

	}
	return nil, errors.New(fmt.Sprintf("User %v was not found", userId))

}