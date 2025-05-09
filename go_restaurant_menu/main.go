package main

import (
	// "go_restaurant_menu/database"
	store "go_restaurant_menu/database"
	"go_restaurant_menu/middleware"
	"go_restaurant_menu/views"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Инициализация базы данных
	db, err := store.InitDB()
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	_ = db // пока не используем

	// Создание нового экземпляра Gin
	r := gin.Default()

	// Определение открытых маршрутов
	r.GET("/", views.HomePage)
	r.POST("/register-user", func(c *gin.Context) {
		views.RegisterUser(c, "admin")
	})
	r.POST("/login", views.Login)

	// Защищённые маршруты для администраторов
	adminGroup := r.Group("/admin")
	// Применяем middleware IsAdmin ко всем маршрутам внутри этого группы
	applyAdminMiddleware(adminGroup)

	// Вложенные маршруты для администраторов
	{
		adminGroup.POST("/restaurants", views.AddRestaurant)
		adminGroup.POST("/menu-items", views.AddMenuItem)
		adminGroup.POST("/add-categories", views.AddCategories)
		adminGroup.GET("/all-restaurants", views.Getrestaurants)
		adminGroup.DELETE("/delete-restaurants", views.DeleteRestaurants)
	}

	// Запуск сервера на порту 8080
	r.Run(":8080")
}

// applyAdminMiddleware автоматически защищает все маршруты внутри группы
func applyAdminMiddleware(group *gin.RouterGroup) {
	group.Use(middleware.IsAdmin()) // Добавляем middleware IsAdmin ко всем маршрутам внутри группы
}
