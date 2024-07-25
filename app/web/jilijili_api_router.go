package web

import (
	"github.com/Ysoding/jilijili/app/controller"
	"github.com/gin-gonic/gin"
)

type JiliJiliAPIRouter struct {
	pingController *controller.PingController
}

func NewJiliJiliAPIRouter(
	pingController *controller.PingController,
) *JiliJiliAPIRouter {
	return &JiliJiliAPIRouter{
		pingController: pingController,
	}
}

func (j *JiliJiliAPIRouter) RegisterPingAPI(r *gin.RouterGroup) {
	r.GET("/ping", j.pingController.HandlePing)
}
