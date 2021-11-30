package entity

type Novel struct {
	Title   string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Context string `json:"context" binding:"max=20"`
	Url     string `json:"url" binding:"required,url"`
	Author  Person `json:"author" binding:"required"`
}
