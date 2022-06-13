package controller

import (
	"github.com/gorilla/mux"
	"log"
	"main/service"
	"net/http"
)

type Controller struct {
	Service service.Service
}

func NewController(service service.Service) *Controller {
	log.Println("controller")
	return &Controller{Service: service}
}

func (c *Controller) GetItem(writer http.ResponseWriter, request *http.Request) {
	result := c.Service.GetItem()
	SendResponse(result.ResponseBody, result.ResponseCode, writer)
}

func (c *Controller) GetItemId(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	result := c.Service.GetItemId(vars["id"])
	SendResponse(result.ResponseBody, result.ResponseCode, writer)
}

func (c *Controller) AddItem(writer http.ResponseWriter, request *http.Request) {
	var item service.Item
	err := DecodeResponse(&item, request)
	if err != nil {
		SendResponse("Request body Error", 400, writer)
		return
	}
	result := c.Service.AddItem(item)
	SendResponse(result.ResponseBody, result.ResponseCode, writer)
}

func (c *Controller) DeleteItemId(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	result := c.Service.DeleteItemId(vars["id"])
	SendResponse(result.ResponseBody, result.ResponseCode, writer)
}

func (c *Controller) UpdateItemId(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	var item service.Item
	err := DecodeResponse(&item, request)
	if err != nil {
		SendResponse("Request body Error", 400, writer)
		return
	}
	result := c.Service.UpdateItemId(vars["id"], item)
	SendResponse(result.ResponseBody, result.ResponseCode, writer)
}

func (c *Controller) AddSampleItem(writer http.ResponseWriter, request *http.Request) {
	var item service.Item
	err := DecodeResponse(&item, request)
	if err != nil {
		SendResponse("Request body Error", 400, writer)
		return
	}
	result := c.Service.AddSampleItem()
	SendResponse(result.ResponseBody, result.ResponseCode, writer)
}
