package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PingController struct {
	log *zap.Logger
}

func NewPingController(log *zap.Logger) *PingController {
	return &PingController{log: log}
}

func (p *PingController) HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (p *PingController) RegisterRoutes(r *gin.RouterGroup) {
	g := r.Group("")

	g.GET("/ping", p.HandlePing)
}
