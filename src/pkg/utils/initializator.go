package utils

var iUserUtil IUserUtil

func InitializeUtils() {
	iUserUtil = NewUserUtil()
}

func UserUtil() IUserUtil {
	return iUserUtil
}
