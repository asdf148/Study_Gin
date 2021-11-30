package main

import (
	"io"
	"os"

	"asdf148.com/Study_Gin/controller"
	"asdf148.com/Study_Gin/middleware"
	"asdf148.com/Study_Gin/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	novelService    service.NovelService       = service.New()
	novelController controller.NovelController = controller.New(novelService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(),
		middleware.BasicAuth(), gindump.Dump())

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, novelController.FindAll())
	})

	server.POST("/novel", func(ctx *gin.Context) {
		ctx.JSON(200, novelController.Save(ctx))
	})

	server.Run() // listen and serve on 0.0.0.0:8080
}
