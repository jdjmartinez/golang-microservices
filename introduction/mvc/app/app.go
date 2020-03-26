package app

import (
	"github.com/jdj.martinez/golang-microservices/introduction/mvc/controllers"
	"net/http"
)

func StartApp()  {
	http.HandleFunc("/user", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}