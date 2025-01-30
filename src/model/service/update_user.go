package service

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/logger"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/rest_err"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdatedUser(
	userId string, 
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser model",
		zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call reppository",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info(
		"UpdateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))
	return nil
}