package controller

import (
	"encoding/json"
	"github.com/jovanitarnowski/golang-mvc/mvc/service"
	"github.com/jovanitarnowski/golang-mvc/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := service.GetUser(userId)
	if err != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
