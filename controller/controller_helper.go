package controller

import (
	"encoding/json"
	"net/http"
)

func SendResponse(i interface{}, status int, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(i)
}

func DecodeResponse(i interface{}, request *http.Request) error {
	return json.NewDecoder(request.Body).Decode(i)
}
