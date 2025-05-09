package store

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"go_restaurant_menu/models"
	"os"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", 
	os.Getenv("DB_HOST"), 
	os.Getenv("DB_USER"), 
	os.Getenv("DB_PASSWORD"), 
	os.Getenv("DB_NAME"), 
	os.Getenv("DB_PORT"), 
	os.Getenv("DB_SSLMODE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	// Автоматическая миграция моделей
	err = db.AutoMigrate(&models.Restaurant{},
		                &models.Category{}, 
						&models.MenuItem{},
						&models.User{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	return db, nil
} 