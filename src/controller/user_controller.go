package controller

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByID(C *gin.Context)
	FindUserByEmail(C *gin.Context)

	DeleteUser(C *gin.Context)
	CreateUser(C *gin.Context)
	UpdatedUser(C *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}