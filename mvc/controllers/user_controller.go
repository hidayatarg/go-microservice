package controllers

import (
	"encoding/json"
	"github.com/hidayatarg/go-microservice/mvc/utils"

	//"github.com/hidayatarg/go-microservice/mvc/domain"
	"github.com/hidayatarg/go-microservice/mvc/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		// return bad request
		apiErr := &utils.ApplicationError {
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// Handle the error and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))

		return
	}
	// otherwise
	// return user to the client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
