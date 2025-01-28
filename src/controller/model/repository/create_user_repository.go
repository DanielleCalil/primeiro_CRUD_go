package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/logger"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/rest_err"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
	_ "github.com/go-sql-driver/mysql"
)

const (
	SQL_USER_DB="SQL_USER_DB"
)

type userRepository struct {
	databaseConnection *sql.DB
}

func (ur *userRepository) CreateUser (
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser repository")

	tableName := os.Getenv(SQL_USER_DB)

	if tableName == "" {
		tableName = "users"
	}

	value, err := userDomain.GetJSONValue() 
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	name, email, password := value["name"].(string), value["email"].(string), value["password"].(string)

	query := fmt.Sprintf("INSERT INTO %s (name, email, password) VALUES (?, ?, ?)", tableName)

	result, err := ur.databaseConnection.Exec(query, name, email, password)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetID(fmt.Sprintf("%d", lastInsertID))

	return userDomain, nil
}