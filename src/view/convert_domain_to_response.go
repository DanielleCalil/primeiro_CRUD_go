package view

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/controller/model/response"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
)

func CovertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}