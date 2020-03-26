package controllers

import (
	"encoding/json"
	"github.com/jdj.martinez/golang-microservices/introduction/mvc/services"
	"github.com/jdj.martinez/golang-microservices/introduction/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := (strconv.ParseInt(userIdParam, 10, 64))

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusNotFound,
			Code:       "bad_request",
		}
		jsoValue, _ := json.Marshal(apiErr)
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsoValue)
		// Just return the Bad Request.
		return
	}

	user, apiErr := services.GetUser(userId)

	if apiErr != nil {
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(apiErr.StatusCode)
		jsoValue, _ := json.Marshal(apiErr)
		resp.Write(jsoValue)
		//Handle err and return
		return
	}

	// return user to client
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	jsoValue, _ := json.Marshal(user)
	resp.Write(jsoValue)
}