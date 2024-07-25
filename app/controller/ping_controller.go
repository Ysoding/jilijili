package controller

import (
	"github.com/Ysoding/jilijili/app/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PingController struct {
	s   *service.TestService
	log *zap.Logger
}

func NewPingController(s *service.TestService, log *zap.Logger) *PingController {
	return &PingController{s: s, log: log}
}

func (p *PingController) HandlePing(c *gin.Context) {
	p.log.Info("test")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
