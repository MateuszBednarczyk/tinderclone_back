package services

var iAccountMaker IAccountMaker
var iAuthorizer IAuthorizer
var iTokenizer IJwtTokenizer

func InitializeServices() {
	iAccountMaker = NewAccountMaker()
	iAuthorizer = NewAuthorizer()
	iTokenizer = NewJwtTokenizer()
}

func RegisterService() IAccountMaker {
	return iAccountMaker
}

func LoginService() IAuthorizer {
	return iAuthorizer
}

func JwtService() IJwtTokenizer {
	return iTokenizer
}
