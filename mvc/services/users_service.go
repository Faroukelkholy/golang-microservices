package services

import (
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
	"log"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	println("services.GetUser")
	//domain.IUserDao = &domain.UserDao{}
	user, err := domain.IUserDao.GetUser(userId)
	log.Println("services.GetUser user", user)
	log.Println("services.GetUser err", err)
	if err != nil {
		println("error")
		return nil, err
	}
	return user, nil
}
