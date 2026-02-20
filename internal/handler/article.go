package handler

import (
	"go-blog/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type articleService interface {
	AllArticles() ([]*model.Article, error)
	ArticleById(id int) (*model.Article, error)
	CreateArticle(newArticle *model.Article) error
	UpdateArticle(id int, updateArticle *model.Article) error
	DeleteArticle(id int) error
}

type articleHandler struct {
	service articleService
	log     *zap.SugaredLogger
}

func NewArticleHandler(service articleService, log *zap.SugaredLogger) *articleHandler {
	return &articleHandler{service: service, log: log}
}

func (h *articleHandler) AllArticles(c *gin.Context) {
	articles, err := h.service.AllArticles()
	if err != nil {
		h.log.Errorf("Ошибка получения всех статей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func (h *articleHandler) ArticleById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Errorf("Ошибка преобразования ID в число: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат ID",
		})
		return
	}

	article, err := h.service.ArticleById(id)
	if err != nil {
		h.log.Errorf("Ошибка получения статьи по ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

func (h *articleHandler) CreateArticle(c *gin.Context) {
	var req model.ArticleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка создания статьи: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат данных",
		})
		return
	}

	newArticle := model.Article{
		Title:   req.Title,
		Content: req.Content,
	}

	err := h.service.CreateArticle(&newArticle)
	if err != nil {
		h.log.Errorf("Ошибка создания статьи: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"article": newArticle,
	})
}

func (h *articleHandler) UpdateArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Errorf("Ошибка преобразования ID в число: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат ID",
		})
		return
	}

	var req model.ArticleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка обновления статьи: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат данных",
		})
		return
	}

	newArticle := model.Article{
		Title:   req.Title,
		Content: req.Content,
	}

	err = h.service.UpdateArticle(id, &newArticle)
	if err != nil {
		h.log.Errorf("Ошибка обновления статьи: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article": newArticle,
	})
}

func (h *articleHandler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Errorf("Ошибка преобразования ID в число: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат ID",
		})
		return
	}

	err = h.service.DeleteArticle(id)
	if err != nil {
		h.log.Errorf("Ошибка удаления статьи: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Статья успешно удалена",
	})
}
