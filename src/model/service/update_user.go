package service

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/rest_err"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
)

func (*userDomainService) UpdatedUser(
	userId string, 
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}