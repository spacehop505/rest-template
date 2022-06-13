package service

import (
	"net/http"
	"sync"
)

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	items map[string]Item
	mux   *sync.Mutex
}

func NewService() *Service {
	return &Service{items: make(map[string]Item, 0)}
}

func (s *Service) GetItem() Response {
	return ServiceResponse(http.StatusOK, s.items)
}

func (s *Service) GetItemId(id string) Response {
	if item, found := s.items[id]; found {
		return ServiceResponse(http.StatusOK, item)
	}
	return ServiceResponse(http.StatusNotFound, "not found")
}

func (s *Service) AddItem(item Item) Response {
	if _, found := s.items[item.Id]; !found {
		s.items[item.Id] = item
		return ServiceResponse(http.StatusCreated, "created item")
	}
	return ServiceResponse(http.StatusBadRequest, "item already exists")
}

func (s *Service) AddSampleItem() Response {
	s.items["a1"] = Item{Id: "a1", Name: "apple"}
	s.items["a2"] = Item{Id: "a2", Name: "orange"}
	s.items["a3"] = Item{Id: "a3", Name: "pear"}
	return ServiceResponse(http.StatusCreated, "created item")
}

func (s *Service) UpdateItemId(id string, item Item) Response {
	if _, found := s.items[id]; found {
		s.items[item.Id] = item
		return ServiceResponse(http.StatusOK, "updated item")
	}
	return ServiceResponse(http.StatusNotFound, "not found")
}

func (s *Service) DeleteItemId(id string) Response {
	if _, found := s.items[id]; found {
		delete(s.items, id)
		return ServiceResponse(http.StatusOK, "removed item")
	}
	return ServiceResponse(http.StatusNotFound, "not found")
}
