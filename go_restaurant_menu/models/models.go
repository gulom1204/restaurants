package models

// import (
// 	"gorm.io/gorm"
// )

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`           // Никогда не возвращай пароль клиенту
	Role     string `json:"role"`        // "admin" или "client"
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Restaurant struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Address      string `json:"address"`
	Phone        string `gorm:"unique" json:"phone"`
	Email        string `gorm:"unique" json:"email"`
	WorkingHours string `json:"working_hours"`
}

type Category struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MenuItem struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	CategoryID  uint    `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	IsAvailable bool    `json:"is_available"`
} 