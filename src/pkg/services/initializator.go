package services

import "tinderclone_back/src/pkg/stores"

var iAccountMaker IAccountMaker
var iAuthenticator IAuthenticator
var iTokenizer IJwtTokenizer
var iAccounter IAccounter
var iCountrier ICountrier

func InitializeServices() {
	iAccountMaker = NewAccountMaker(stores.UserStore(), stores.CountryStore(), stores.CityStore())
	iAuthenticator = NewAuthenticator(stores.UserStore())
	iTokenizer = NewJwtTokenizer()
	iAccounter = NewAccounter(stores.UserStore())
	iCountrier = NewCountrier(stores.CountryStore())
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
