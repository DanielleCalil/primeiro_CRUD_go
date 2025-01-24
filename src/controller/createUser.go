package controller

import (
	"net/http"

	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/logger"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/configuration/validation"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/controller/model/request"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/model"
	"github.com/DanielleCalil/primeiro_CRUD_go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init createUser function",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created secessfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.CovertDomainToResponse(
		domain,
	))
}
