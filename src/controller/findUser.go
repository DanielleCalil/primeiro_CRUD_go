package controller

import (
	"net/http"
	"net/mail"

	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/logger"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/rest_err"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserByID controller",
		zap.String("journey", "FindUserByID"),
	)

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "FindUserByID"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"Invalid user ID",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call FindUserByID services",
			err,
			zap.String("journey", "FindUserByID"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
	)
	c.JSON(http.StatusOK, view.CovertDomainToResponse(
		userDomain,
	))
}


func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller",
	zap.String("journey", "FindUserByEmail"),
)

userEmail := c.Param("userEmail")

if _, err := mail.ParseAddress(userEmail); err != nil {
	logger.Error("Error trying to validate userEmail",
		err,
		zap.String("journey", "FindUserByEmail"),
	)
	errorMessage := rest_err.NewBadRequestError(
		"Invalid user Email",
	)

	c.JSON(errorMessage.Code, errorMessage)
	return
}

userDomain, err := uc.service.FindUserByEmailServices(userEmail)
if err != nil {
	logger.Error("Error trying to call FindUserByEmail services",
		err,
		zap.String("journey", "FindUserByEmail"),
	)
	c.JSON(err.Code, err)
	return
}

logger.Info("FindUserByEmail controller executed successfully",
	zap.String("journey", "findUserByEmail"),
)
c.JSON(http.StatusOK, view.CovertDomainToResponse(
	userDomain,
))
}
