package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null;type:text" json:"content"`
}

type ArticleCreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type ArticleUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a ArticleCreateRequest) ToArticle() Article {
	return Article{
		Title:   a.Title,
		Content: a.Content,
	}
}

func (a ArticleUpdateRequest) ToArticle() Article {
	return Article{
		Title:   a.Title,
		Content: a.Content,
	}
}
