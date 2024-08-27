package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	UserID uint `json:"user_id" gorm:"not null"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	CategoryID uint `json:"category_id" gorm:"not null"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
	Comments []Comment `json:"comments" gorm:"foreignKey:BlogID"`
	Status  string `json:"status" gorm:"type:enum('draft', 'published', 'archived');default:'published'"`
}

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Blogs []Blog `json:"blogs" gorm:"foreignKey:CategoryID"`
	Status string `json:"status" gorm:"type:enum('active', 'inactive');default:'active'"`
}

type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"not null"`
	UserID uint `json:"user_id" gorm:"not null"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	BlogID uint `json:"blog_id" gorm:"not null"`
	Blog Blog `json:"blog" gorm:"foreignKey:BlogID"`
	Status string `json:"status" gorm:"type:enum('active', 'inactive');default:'active'"`
}