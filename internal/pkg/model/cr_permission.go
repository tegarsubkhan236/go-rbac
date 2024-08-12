package model

import "time"

type CrPermission struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(255);NOT NULL" json:"name"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"type:timestamp" json:"deleted_at,omitempty"`
}

func (receiver CrPermission) ToResponse() *CrPermission {
	return &CrPermission{}
}