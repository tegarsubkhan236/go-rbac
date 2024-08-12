package model

import (
	"time"
)

type CrRole struct {
	ID          int             `gorm:"primary_key"`
	Name        string          `gorm:"type:varchar(255);not null"`
	Permissions []*CrPermission `gorm:"many2many:cr_role_permissions;"`
	CreatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP() NOT NULL;type:datetime"`
	UpdatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP() NOT NULL;type:timestamp"`
	DeletedAt   *time.Time      `gorm:"type:timestamp;null"`
}
