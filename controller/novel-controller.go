package controller

import (
	"net/http"

	"asdf148.com/Study_Gin/entity"
	"asdf148.com/Study_Gin/service"
	"asdf148.com/Study_Gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NovelController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Novel
	ShowAll(ctx *gin.Context)
}

type novelController struct {
	service service.NovelService
}

var validate *validator.Validate

func New(service service.NovelService) NovelController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &novelController{
		service: service,
	}
}

func (c *novelController) Save(ctx *gin.Context) error {
	var novel entity.Novel
	err := ctx.ShouldBindJSON(&novel)
	if err != nil {
		return err
	}
	err = validate.Struct(novel)
	if err != nil {
		return err
	}
	c.service.Save(novel)
	return nil
}

func (c *novelController) FindAll() []entity.Novel {
	return c.service.FindAll()
}

func (c *novelController) ShowAll(ctx *gin.Context) {
	novels := c.service.FindAll()
	data := gin.H{
		"title":  "Novel Page",
		"novels": novels,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
