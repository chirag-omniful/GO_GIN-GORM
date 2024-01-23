package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Name string `json:"name"`
	Body string `json:"body"`
}

// jo gorm.model likha h us se yeh sab fields add hongi
// gorm.Model definition
//type Model struct {
//	ID        uint           `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//}
