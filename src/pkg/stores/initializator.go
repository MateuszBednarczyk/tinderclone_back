package stores

import "gorm.io/gorm"

var iUserStore IUserStore
var iCountryStore ICountryStore
var iCityStore ICityStore

func InitializeStores(db *gorm.DB) {
	iUserStore = NewUserStore(db)
	iCountryStore = NewCountryStore(db)
	iCityStore = NewCityStore(db)
}

func UserStore() IUserStore {
	return iUserStore
}

func CountryStore() ICountryStore {
	return iCountryStore
}

func CityStore() ICityStore {
	return iCityStore
}
