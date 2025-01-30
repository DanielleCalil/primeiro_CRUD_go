package controller

import (
	"net/http"
	"strings"

	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/logger"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/validation"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/controller/model/request"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdatedUser(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest

	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&userRequest); err != nil ||
	strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

		domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)
	err := uc.service.UpdatedUser(userId, domain); 
	if err != nil {
		logger.Error(
			"Error trying to call UpdateUser service", 
			err,
			zap.String("journey", "UpdateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "UpdateUser"))

	c.Status(http.StatusOK)
}
