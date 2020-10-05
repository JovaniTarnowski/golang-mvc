package app

import (
	"github.com/jovanitarnowski/golang-mvc/mvc/controller"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controller.GetUser)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
