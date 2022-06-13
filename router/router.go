package router

import (
	"github.com/gorilla/mux"
	"log"
	"main/controller"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(controller controller.Controller) *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes(controller) {
		log.Println("route initialized:", route.Method, route.Pattern, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

func routes(controller controller.Controller) Routes {
	return Routes{
		{Name: "GetItem", Method: "GET", Pattern: "/", HandlerFunc: controller.GetItem},
		{Name: "GetItemId", Method: "GET", Pattern: "/{id}", HandlerFunc: controller.GetItemId},
		{Name: "AddItem", Method: "POST", Pattern: "/", HandlerFunc: controller.AddItem},
		{Name: "AddSampleItem", Method: "POST", Pattern: "/sample", HandlerFunc: controller.AddSampleItem},
		{Name: "UpdateItemId", Method: "PUT", Pattern: "/{id}", HandlerFunc: controller.UpdateItemId},
		{Name: "DeleteItemId", Method: "DELETE", Pattern: "/{id}", HandlerFunc: controller.DeleteItemId},
	}
}
