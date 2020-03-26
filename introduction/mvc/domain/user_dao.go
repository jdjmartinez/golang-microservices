package domain

import (
	"fmt"
	"github.com/jdj.martinez/golang-microservices/introduction/mvc/utils"
	"net/http"
)

var(
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Fede", LastName: "Leon", Email: "ejemploLeon@ejemplo.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError)  {

	if user := users[userId]; user != nil{
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
