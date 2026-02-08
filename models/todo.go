package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Username  string         `gorm:"unique;not null" json:"username" binding:"required"`
	Password  string         `gorm:"not null" json:"password" binding:"required"`
}

// Todo 待办事项模型
type Todo struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `gorm:"not null" json:"title" binding:"required"`
	Description string         `json:"description"`
	IsCompleted bool           `gorm:"default:false" json:"is_completed"`
	UserID      uint           `gorm:"not null" json:"user_id"`
}
