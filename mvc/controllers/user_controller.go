package controllers

import (
	"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 10)
	if err != nil {
		// return bad request
		return
	}

	user, err :=
}
