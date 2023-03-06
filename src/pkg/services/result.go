package services

type Result struct {
	Message string
	Code    int
	Content []interface{}
}

func CreateServiceResult(message string, code int, content []interface{}) *Result {
	return &Result{
		Message: message,
		Code:    code,
		Content: content,
	}
}
