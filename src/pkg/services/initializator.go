package services

var iAccountMaker IAccountMaker
var iAuthenticator IAuthenticator
var iTokenizer IJwtTokenizer
var iAccounter IAccounter
var iCountrier ICountrier

func InitializeServices() {
	iAccountMaker = NewAccountMaker()
	iAuthenticator = NewAuthenticator()
	iTokenizer = NewJwtTokenizer()
	iAccounter = NewAccounter()
	iCountrier = NewCountrier()
}

func AccountMaker() IAccountMaker {
	return iAccountMaker
}

func Authenticator() IAuthenticator {
	return iAuthenticator
}

func Tokenizer() IJwtTokenizer {
	return iTokenizer
}

func Accounter() IAccounter {
	return iAccounter
}

func Countrier() ICountrier {
	return iCountrier
}
