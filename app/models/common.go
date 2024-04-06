package models

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type CreatedAt struct {
	CreatedAt time.Time `json:"created_at"`
}

type UpdatedAt struct {
	UpdatedAt time.Time `json:"updated_at"`
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
