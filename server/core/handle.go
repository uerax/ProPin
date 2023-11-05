package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uerax/propin/server/pin"
)

func Report(ctx *gin.Context) {
	var status Status
	if ctx.ShouldBind(&status) == nil {
		fmt.Println("Add: ",status.Name,":",ctx.RemoteIP())
		pin.Add(status.Name, ctx.RemoteIP())
	}
}

func Info(ctx *gin.Context) {
	msg := pin.Info()
	go http.Get(tgApi+msg)
	ctx.JSON(200, gin.H{
		"message": "已发送到Telegram",
	})
}

func ToTg(ctx *gin.Context) {
	msg := pin.Out()
	fmt.Println(msg)
	go http.Get(tgApi+msg)
	ctx.JSON(200, gin.H{
		"message": "已发送到Telegram",
	})
}