package router

import (
	"fmt"
	"go-blog/internal/config"
	"go-blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

type userHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type articleHandler interface {
	AllArticles(c *gin.Context)
	ArticleById(c *gin.Context)
	CreateArticle(c *gin.Context)
	UpdateArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
}

func NewRouter(cfg *config.Config, userHandler userHandler, articleHandler articleHandler) (*gin.Engine, error) {
	var r *gin.Engine
	switch cfg.Environment {
	case "dev":
		r = gin.Default()
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery(), gin.Logger())
	default:
		return nil, fmt.Errorf("Не верно указано окружение %s", cfg.Environment)
	}
	r.Use(middleware.CORSMiddleware(cfg.CORS))
	api := r.Group("/api")

	api.POST("/users/register", userHandler.Register)
	api.POST("/users/login", userHandler.Login)

	api.GET("/articles", articleHandler.AllArticles)
	api.GET("/articles/:id", articleHandler.ArticleById)

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.POST("/articles", articleHandler.CreateArticle)
	admin.PUT("/articles/:id", articleHandler.UpdateArticle)
	admin.DELETE("/articles/:id", articleHandler.DeleteArticle)

	return r, nil
}
