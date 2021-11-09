package main

import (
	"github.com/Elder6Driver/grayReleaseProject/service"
	"github.com/gin-gonic/gin"
)

func customizeouter(r *gin.Engine) {
	r.GET("/ping", service.Pong)
	r.GET("/judge", service.Hit)
}
