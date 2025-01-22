package routes

import (
	"github.com/DanielleCalil/primeiro_CRUD_go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdatedUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}