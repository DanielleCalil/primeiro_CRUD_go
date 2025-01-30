package main

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/controller"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model/repository"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	//Init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
