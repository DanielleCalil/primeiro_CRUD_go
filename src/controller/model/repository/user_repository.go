package repository

import (
	"database/sql"

	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/rest_err"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
)

func NewUserRepository(
	datebase *sql.DB,
) UserRepository {
	return &userRepository{
		datebase,
	}
}
type userRepository struct {
	datebaseConnection *sql.DB
}
type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}