package converter

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model/repository/entity"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())

	return domain
}