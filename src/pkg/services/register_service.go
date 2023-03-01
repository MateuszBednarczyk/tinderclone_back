package services

type IRegisterService interface {
	RegisterUser() *Result
}

type registerService struct {
}

func NewRegisterService() *registerService {
	return &registerService{}
}

func (s *registerService) RegisterUser() *Result {
	return &Result{}
}
