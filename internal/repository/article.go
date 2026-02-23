package repository

import (
	"go-blog/internal/model"

	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db: db}
}

func (s articleRepository) GetArticles() ([]model.Article, error) {
	var articles []model.Article
	if err := s.db.Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (s articleRepository) GetArticle(id int) (model.Article, error) {
	var article model.Article
	if err := s.db.First(&article, id).Error; err != nil {
		return model.Article{}, err
	}

	return article, nil
}

func (s articleRepository) CreateArticle(newArticle model.Article) error {
	return s.db.Create(&newArticle).Error
}

func (s articleRepository) UpdateArticle(id int, updateArticle model.Article) error {
	var article model.Article
	if err := s.db.First(&article, id).Error; err != nil {
		return err
	}

	if err := s.db.Model(&article).Updates(updateArticle).Error; err != nil {
		return err
	}

	return nil
}

func (s articleRepository) DeleteArticle(id int) error {
	var article model.Article
	if err := s.db.First(&article, id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&article).Error; err != nil {
		return err
	}

	return nil
}
