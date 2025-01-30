package repository

import (
	"context"
	"os"

	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/logger"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/rest_err"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser repository",
		zap.String("journey", "updateUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to updateUser",
			err,
			zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"updateUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	return nil
}
