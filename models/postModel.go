package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

// jo gorm.model likha h us se yeh sab fields add hongi
// gorm.Model definition
//type Model struct {
//	ID        uint           `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//}
