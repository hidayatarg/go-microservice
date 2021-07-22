package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hidayatarg/go-microservice/mvc/utils"
	"github.com/hidayatarg/go-microservice/mvc/services"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		// return bad request
		apiErr := &utils.ApplicationError {
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad request",
		}

		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	// valid userId
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// Handle the error and return to the client
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	// otherwise
	// return user to the client
	c.JSON(http.StatusOK, user)
}
