package services

var iRegisterService IRegisterService

func InitializeServices() {
	iRegisterService = NewRegisterService()
}

func RegisterService() IRegisterService {
	return iRegisterService
}
