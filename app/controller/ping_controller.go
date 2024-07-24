package controller

import (
	"github.com/Ysoding/jilijili/app/service"
	"github.com/gin-gonic/gin"
)

type PingController struct {
	s *service.TestService
}

func NewPingController(s *service.TestService) *PingController {
	return &PingController{s: s}
}

func (p *PingController) HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
