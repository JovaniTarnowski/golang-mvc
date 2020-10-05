package service

import (
	"github.com/jovanitarnowski/golang-mvc/mvc/domain"
	"github.com/jovanitarnowski/golang-mvc/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
