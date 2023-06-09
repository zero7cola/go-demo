package user

import "go1/pkg/database"

func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)

	return count > 0
}

func IsPhoneExist(phone string) bool {

	var count int64
	//database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)

	return count > 0
}
