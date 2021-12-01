package main

import (
	"io"
	"net/http"
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

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middleware.Logger(),
		middleware.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/posts", func(ctx *gin.Context) {
			ctx.JSON(200, novelController.FindAll())
		})

		apiRoutes.POST("/novel", func(ctx *gin.Context) {
			err := novelController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "novel Input is Valid!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/novels", novelController.ShowAll)
	}

	server.Run() // listen and serve on 0.0.0.0:8080
}
