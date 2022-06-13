package service

type Response struct {
	ResponseCode int
	ResponseBody interface{}
}

func ServiceResponse(responseCode int, responseBody interface{}) Response {
	return Response{ResponseCode: responseCode, ResponseBody: responseBody}
}
