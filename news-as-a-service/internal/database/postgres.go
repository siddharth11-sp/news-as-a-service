package database

import (
	"fmt"
	"log"

	"github.com/siddharth11-sp/news-service/internal/config"
	"github.com/siddharth11-sp/news-service/internal/news"

	"github.com/siddharth11-sp/news-service/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect db: %w", err)
	}

	if err := pingDB(db); err != nil {
		return nil, err
	}

	if err := autoMigrate(db); err != nil {
		return nil, err
	}

	log.Println("postgres connected")

	return db, nil
}

func pingDB(db *gorm.DB) error {

	sqlDB, err := db.DB()

	if err != nil {
		return fmt.Errorf("failed getting sql db: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("db ping failed: %w", err)
	}

	return nil
}

func autoMigrate(db *gorm.DB) error {

	return db.AutoMigrate(
		&entity.Entity{},
		&news.NewsArticle{},
	)
}
