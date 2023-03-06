package handlers

import "tinderclone_back/src/pkg/services"

type Response struct {
	Message string
	Content []interface{}
}

func CreateHandlerResponse(r *services.Result) *Response {
	return &Response{
		Message: r.Message,
		Content: r.Content,
	}
}
