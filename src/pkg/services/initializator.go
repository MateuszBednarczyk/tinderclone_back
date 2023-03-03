package services

var iRegisterService IRegisterService
var iLoginService ILoginService

func InitializeServices() {
	iRegisterService = NewRegisterService()
	iLoginService = NewLoginService()
}

func RegisterService() IRegisterService {
	return iRegisterService
}

func LoginService() ILoginService {
	return iLoginService
}
