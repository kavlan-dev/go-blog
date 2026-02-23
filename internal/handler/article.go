package handler

import (
	"go-blog/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type articleService interface {
	AllArticles() ([]model.Article, error)
	ArticleById(id int) (model.Article, error)
	CreateArticle(newArticle model.ArticleCreateRequest) error
	UpdateArticle(id int, updateArticle model.ArticleUpdateRequest) error
	DeleteArticle(id int) error
}

type articleHandler struct {
	s   articleService
	log *zap.SugaredLogger
}

func NewArticleHandler(s articleService, log *zap.SugaredLogger) *articleHandler {
	return &articleHandler{s: s, log: log}
}

func (h articleHandler) AllArticles(c *gin.Context) {
	articles, err := h.s.AllArticles()
	if err != nil {
		h.log.Errorf("Ошибка получения всех статей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func (h articleHandler) ArticleById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Errorf("Ошибка преобразования ID в число: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный формат ID",
		})
		return
	}

	article, err := h.s.ArticleById(id)
	if err != nil {
		h.log.Errorf("Ошибка получения статьи по ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

func (h articleHandler) CreateArticle(c *gin.Context) {
	var req model.ArticleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка создания статьи: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный формат данных",
		})
		return
	}

	err := h.s.CreateArticle(req)
	if err != nil {
		h.log.Errorf("Ошибка создания статьи: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "статья успешно создана",
	})
}

func (h articleHandler) UpdateArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Errorf("Ошибка преобразования ID в число: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный формат ID",
		})
		return
	}

	var req model.ArticleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка обновления статьи: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный формат данных",
		})
		return
	}

	err = h.s.UpdateArticle(id, req)
	if err != nil {
		h.log.Errorf("Ошибка обновления статьи: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "статья успешно обновлена",
	})
}

func (h articleHandler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Errorf("Ошибка преобразования ID в число: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный формат ID",
		})
		return
	}

	err = h.s.DeleteArticle(id)
	if err != nil {
		h.log.Errorf("Ошибка удаления статьи: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "статья успешно удалена",
	})
}
