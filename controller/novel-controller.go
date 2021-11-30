package controller

import (
	"asdf148.com/Study_Gin/entity"
	"asdf148.com/Study_Gin/service"
	"github.com/gin-gonic/gin"
)

type NovelController interface {
	Save(ctx *gin.Context) entity.Novel
	FindAll() []entity.Novel
}

type novelController struct {
	service service.NovelService
}

func New(service service.NovelService) NovelController {
	return &novelController{
		service: service,
	}
}

func (c *novelController) Save(ctx *gin.Context) entity.Novel {
	var novel entity.Novel
	ctx.ShouldBindJSON(&novel)
	c.service.Save(novel)
	return novel
}

func (c *novelController) FindAll() []entity.Novel {
	return c.service.FindAll()
}
