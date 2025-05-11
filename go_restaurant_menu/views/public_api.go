package views

import (
	store "go_restaurant_menu/database"
	"go_restaurant_menu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPublicMenu возвращает публичное меню
func GetPublicMenu(c *gin.Context) {
	db, err := store.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var menuItems []models.MenuItem
	if err := db.Find(&menuItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu items"})
		return
	}

	c.JSON(http.StatusOK, menuItems)
}

// GetPublicCategories возвращает публичные категории
func GetPublicCategories(c *gin.Context) {
	db, err := store.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetPublicRestaurants возвращает публичные рестораны
func GetPublicRestaurants(c *gin.Context) {
	db, err := store.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var restaurants []models.Restaurant
	if err := db.Find(&restaurants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch restaurants"})
		return
	}

	c.JSON(http.StatusOK, restaurants)
} 