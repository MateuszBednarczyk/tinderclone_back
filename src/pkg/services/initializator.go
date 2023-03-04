package services

var iRegisterService IRegisterService
var iLoginService ILoginService
var iJWTService IJWTService

func InitializeServices() {
	iRegisterService = NewRegisterService()
	iLoginService = NewLoginService()
	iJWTService = NewJwtService()
}

func RegisterService() IRegisterService {
	return iRegisterService
}

func LoginService() ILoginService {
	return iLoginService
}

func JwtService() IJWTService {
	return iJWTService
}
