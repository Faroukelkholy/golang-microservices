package services

import (
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	println("services.GetUser")
	//domain.IUserDao = &domain.UserDao{}
	user, err := domain.IUserDao.GetUser(userId)
	if err != nil {
		println("error")
		return nil, err
	}
	return user, nil
}
