package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
