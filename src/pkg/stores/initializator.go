package stores

var iUserStore IUserStore
var iCountryStore ICountryStore
var iCityStore ICityStore

func InitializeStores() {
	iUserStore = NewUserStore()
	iCountryStore = NewCountryStore()
	iCityStore = NewCityStore()
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
