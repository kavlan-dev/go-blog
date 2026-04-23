package app

import (
	"context"
	"go-blog/internal/config"
	"go-blog/internal/handler"
	"go-blog/internal/repository"
	"go-blog/internal/repository/postgres"
	"go-blog/internal/router"
	"go-blog/internal/service"
	"go-blog/internal/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации конфигурации: %v", err)
	}

	logger, err := util.NewLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}
	defer logger.Sync()

	db, err := postgres.New(cfg)
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, logger)

	articleRepository := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepository)
	articleHandler := handler.NewArticleHandler(articleService, logger)

	r, err := router.NewRouter(cfg, userHandler, articleHandler)
	if err != nil {
		log.Fatalf("Ошибка инициализации роутера: %v", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("Сервер запущен")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalln("Ошибка запуска сервера:", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logger.Infoln("Остановка сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("Ошибка graceful shutdown: %v", err)
		return
	}

	logger.Infoln("Сервер успешно завершил работу")
}
