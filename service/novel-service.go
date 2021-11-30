package service

import (
	"asdf148.com/Study_Gin/entity"
)

type NovelService interface {
	Save(entity.Novel) entity.Novel
	FindAll() []entity.Novel
}

type novelService struct {
	novles []entity.Novel
}

func New() NovelService {
	return &novelService{}
}

func (service *novelService) Save(novel entity.Novel) entity.Novel {
	service.novles = append(service.novles, novel)
	return novel
}

func (service *novelService) FindAll() []entity.Novel {
	return service.novles
}
