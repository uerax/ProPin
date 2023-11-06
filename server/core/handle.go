package core

import (
	"fmt"
	"net/http"
	"net/url"

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
	data := url.Values{}
	data.Add("chat_id", chatId)
	data.Add("text", msg)
	go http.PostForm(tgApi, data)
	ctx.JSON(200, gin.H{
		"message": "已发送到Telegram",
	})
}

func ToTg(ctx *gin.Context) {
	msg := pin.Out()
	fmt.Println(msg)
	data := url.Values{}
	data.Add("chat_id", chatId)
	data.Add("text", msg)
	go http.PostForm(tgApi, data)
	
	ctx.JSON(200, gin.H{
		"message": "已发送到Telegram",
	})
}