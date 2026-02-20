package handler

import (
	"go-blog/internal/model"
	"go-blog/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userService interface {
	AuthenticateUser(authUser *model.User) (*model.User, error)
	RegisterUser(newUser *model.User) error
}

type userHandler struct {
	service userService
	log     *zap.SugaredLogger
}

func NewUserHandler(service userService, log *zap.SugaredLogger) *userHandler {
	return &userHandler{service: service, log: log}
}

func (h *userHandler) Register(c *gin.Context) {
	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка в теле запроса регистрации: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верное тело запроса",
		})
		return
	}

	newUser := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := h.service.RegisterUser(&newUser); err != nil {
		h.log.Errorf("Ошибка при создании пользователя: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "не удалось создать пользователя",
		})
		return
	}

	h.log.Debugf("Успешное создание пользователя")
	c.JSON(http.StatusCreated, gin.H{
		"message": "пользователь успешно создан",
	})
}

func (h *userHandler) Login(c *gin.Context) {
	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("Ошибка в теле запроса логина: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "не верное тело запроса",
		})
		return
	}

	newUser := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := h.service.AuthenticateUser(&newUser)
	if err != nil {
		h.log.Errorf("Ошибка авторизации пользователя %s: %v", req.Username, err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "неверное имя пользователя или пароль",
		})
		return
	}

	token, err := util.GenerateJWT(user.ID)
	if err != nil {
		h.log.Errorf("Ошибка генерации токена для пользователя #%d: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "не удалось сгенерировать JWT токен",
		})
		return
	}

	h.log.Debugf("Пользователь #%d успешно вошел", user.ID)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
