package handlers

import "tinderclone_back/src/pkg/service"

type Response struct {
	Message string
	Content []interface{}
}

func CreateResponse(r *service.Result) *Response {
	return &Response{
		Message: r.Message,
		Content: r.Content,
	}
}
