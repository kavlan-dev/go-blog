package service

import "go-blog/internal/model"

type articleStorage interface {
	GetArticles() ([]*model.Article, error)
	GetArticle(id int) (*model.Article, error)
	CreateArticle(newArticle *model.Article) error
	UpdateArticle(id int, updateArticle *model.Article) error
	DeleteArticle(id int) error
}

type articleService struct {
	db articleStorage
}

func NewArticleService(db articleStorage) articleService {
	return articleService{db: db}
}

func (s articleService) AllArticles() ([]*model.Article, error) {
	return s.db.GetArticles()
}

func (s articleService) ArticleById(id int) (*model.Article, error) {
	return s.db.GetArticle(id)
}

func (s articleService) CreateArticle(newArticle *model.Article) error {

	return s.db.CreateArticle(newArticle)
}

func (s articleService) UpdateArticle(id int, updateArticle *model.Article) error {
	return s.db.UpdateArticle(id, updateArticle)
}

func (s articleService) DeleteArticle(id int) error {
	return s.db.DeleteArticle(id)
}
