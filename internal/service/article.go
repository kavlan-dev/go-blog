package service

import "go-blog/internal/model"

type articleRepository interface {
	GetArticles() ([]model.Article, error)
	GetArticle(id int) (model.Article, error)
	CreateArticle(newArticle model.Article) error
	UpdateArticle(id int, updateArticle model.Article) error
	DeleteArticle(id int) error
}

type articleService struct {
	repo articleRepository
}

func NewArticleService(repo articleRepository) *articleService {
	return &articleService{repo: repo}
}

func (s articleService) AllArticles() ([]model.Article, error) {
	return s.repo.GetArticles()
}

func (s articleService) ArticleById(id int) (model.Article, error) {
	return s.repo.GetArticle(id)
}

func (s articleService) CreateArticle(newArticle model.ArticleCreateRequest) error {
	article := newArticle.ToArticle()
	return s.repo.CreateArticle(article)
}

func (s articleService) UpdateArticle(id int, updateArticle model.ArticleUpdateRequest) error {
	article := updateArticle.ToArticle()
	return s.repo.UpdateArticle(id, article)
}

func (s articleService) DeleteArticle(id int) error {
	return s.repo.DeleteArticle(id)
}
