package main

import (
	"main/controller"
	"main/router"
	"main/service"
	"net/http"
)

func main() {
	newService := service.NewService()
	newController := controller.NewController(*newService)
	newRouter := router.NewRouter(*newController)

	srv := &http.Server{
		Handler: newRouter,
		Addr:    "127.0.0.1:8000",
	}

	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
