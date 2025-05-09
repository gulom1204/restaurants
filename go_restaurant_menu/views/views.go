package views

import (
	"fmt"
	store "go_restaurant_menu/database"
	"go_restaurant_menu/models"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func HomePage(c *gin.Context) {
	var restaurants []models.Restaurant
	var categories []models.Category
	var items []models.MenuItem
	db, err := store.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err := db.Find(&restaurants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err := db.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"restaurants": restaurants,
		"categories": categories,
		"items": items,
	})
}

func RegisterUser(c *gin.Context, role string) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат"})
		return
	}

	user.Role = role
	log.Println("Пароль перед хешированием:", user.Password)

	hash, err := hashPassword(strings.TrimSpace(user.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеширования пароля"})
		return
	}
	user.Password = hash
	db, _ := store.InitDB()

	var existing models.User
	if err := db.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email уже зарегистрирован"})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь зарегистрирован"})
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Println("-------------> ", string(hash))
	return string(hash), nil
}

func Login(c *gin.Context) {
    var input models.LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ввод"})
        return
    }

    db, _ := store.InitDB()
    var user models.User
    if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
        return
    }

    // Логируем введенные данные и хеш из базы данных для отладки
    log.Println("Проверка пароля:", input.Password)
    log.Println("Хеш пароля из базы данных:", user.Password)

    // Убираем пробелы с пароля перед сравнением
    inputPassword := strings.TrimSpace(input.Password)

    // Проверяем пароль
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword)); err != nil {
        log.Println("Ошибка сравнения паролей:", err)  // Логируем ошибку
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
        return
    }

    // Создание токена
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "role":    user.Role,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func AddRestaurant(c *gin.Context) {
	var input models.Restaurant
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ввод"})
		return
	}

	// Инициализируем базу данных
	db, _ := store.InitDB()

	// Добавляем ресторан в базу данных
	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func AddCategories(c *gin.Context) {
	var category models.Category

	// Привязка данных из JSON тела запроса
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// Инициализация базы данных
	db, err := store.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Создание категории в базе данных
	if err := db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем успешный ответ с добавленной категорией
	c.JSON(http.StatusOK, gin.H{"category": category})
}

func AddMenuItem(c *gin.Context) {
	var input models.MenuItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ввод"})
		return
	}

	// Инициализация базы данных
	db, _ := store.InitDB()

	// Проверка, существует ли категория с таким ID
	var category models.Category
	if err := db.First(&category, input.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Категория с таким ID не существует"})
		return
	}

	// Добавление элемента меню в базу данных
	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func Getrestaurants(c *gin.Context) {
    var restaurants []models.Restaurant
    db, err := store.InitDB()

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка подключения к базе данных"})
        return
    }

    // Поиск всех ресторанов
    if err := db.Find(&restaurants).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных ресторанов"})
        return
    }

    // Если рестораны не найдены
    if len(restaurants) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Рестораны не найдены"})
        return
    }

    // Возвращаем найденные рестораны
    c.JSON(http.StatusOK, gin.H{"restaurants": restaurants})
}

func DeleteRestaurants(c *gin.Context) {
	db, err := store.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Удаляем все записи
	if err := db.Where("1 = 1").Delete(&models.Restaurant{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении данных"})
		return
	}

	// Сброс автоинкремента (если используете PostgreSQL)
	if err := db.Exec("ALTER SEQUENCE restaurants_id_seq RESTART WITH 1").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сброса последовательности"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Все рестораны удалены и последовательность сброшена"})
}
