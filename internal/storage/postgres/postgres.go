package postgres

import (
	"fmt"
	"go-blog/internal/config"
	"go-blog/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func NewStorage(cfg *config.Config) (*storage, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}, &model.Article{}); err != nil {
		return nil, err
	}

	return &storage{db: db}, nil
}
