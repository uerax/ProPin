package core

import "github.com/gin-gonic/gin"

var (
	tgApi string
	chatId string
)

func Start(api, id string) {
	s := gin.Default()
	tgApi = api
	chatId = id
	s.POST("/report", Report)
	s.GET("/info", Info)
	s.GET("/status", ToTg)
	s.Run(":4499")
}

