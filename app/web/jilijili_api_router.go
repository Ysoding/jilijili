package web

import (
	"github.com/Ysoding/jilijili/app/controller"
	"github.com/gin-gonic/gin"
)

type JiliJiliAPIRouter struct {
	pingController *controller.PingController
	userController *controller.UserController
}

func NewJiliJiliAPIRouter(
	pingController *controller.PingController,
	userController *controller.UserController,
) *JiliJiliAPIRouter {
	return &JiliJiliAPIRouter{
		pingController: pingController,
		userController: userController,
	}
}

func (j *JiliJiliAPIRouter) RegisterRoutes(server *gin.Engine) {
	g := server.Group("/v1")
	j.pingController.RegisterRoutes(g)
	j.userController.RegisterRoutes(g)
	// r.GET("/v1/ping", j.pingController.HandlePing)
}
