package service

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/logger"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/rest_err"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID services",
		zap.String("journey", "FindUserByID"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail services",
		zap.String("journey", "FindUserByEmail"))

	return ud.userRepository.FindUserByEmail(id)
}
