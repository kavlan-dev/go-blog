package app

import (
	"context"
	"go-blog/internal/config"
	"go-blog/internal/handler"
	"go-blog/internal/router"
	"go-blog/internal/service"
	"go-blog/internal/storage/postgres"
	"go-blog/internal/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	db, err := postgres.NewStorage(cfg)
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService, logger)

	articleService := service.NewArticleService(db)
	articleHandler := handler.NewArticleHandler(articleService, logger)

	r, err := router.NewRouter(cfg, userHandler, articleHandler)
	if err != nil {
		log.Fatalf("Ошибка инициализации роутера: %v", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	logger.Info("Сервер запущен")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Info("Остановка...")
		cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf("Ошибка остановки сервера %v:", err)
		}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
