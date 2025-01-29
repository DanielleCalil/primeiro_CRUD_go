package converter

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/controller/model/repository/entity"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}