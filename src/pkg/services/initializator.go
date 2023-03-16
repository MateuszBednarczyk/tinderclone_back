package services

import "tinderclone_back/src/pkg/stores"

var iAccountMaker IAccountMaker
var iAuthenticator IAuthenticator
var iTokenizer IJwtTokenizer
var iAccounter IAccounter
var iCountrier ICountrier
var iCitier ICitier
var iPermitter IPermitter

func InitializeServices() {
	iAccountMaker = NewAccountMaker(stores.UserStore(), stores.CountryStore(), stores.CityStore())
	iAuthenticator = NewAuthenticator(stores.UserStore())
	iTokenizer = NewJwtTokenizer()
	iAccounter = NewAccounter(stores.UserStore())
	iCountrier = NewCountrier(stores.CountryStore())
	iCitier = NewCitier(stores.CityStore(), stores.CountryStore())
	iPermitter = NewPermitter(stores.UserStore())
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

func Citier() ICitier {
	return iCitier
}

func Permitter() IPermitter {
	return iPermitter
}
