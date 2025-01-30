package service

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/rest_err"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model/repository"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDServices(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailServices(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdatedUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
