package postgres

import (
	"fmt"
	"go-blog/internal/config"
	"go-blog/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}, &model.Article{}); err != nil {
		return nil, err
	}

	return db, nil
}
