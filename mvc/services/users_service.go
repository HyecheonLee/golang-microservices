package services

import (
	"github.com/HyecheonLee/golang-microservices/mvc/domain"
	"github.com/HyecheonLee/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
