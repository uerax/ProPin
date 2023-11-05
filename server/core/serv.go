package core

import "github.com/gin-gonic/gin"

var tgApi string

func Start(api string) {
	s := gin.Default()
	tgApi = api
	s.POST("/report", Report)
	s.GET("/info", Info)
	s.GET("/status", ToTg)
	s.Run(":4499")
}

