package model

import "time"

type CrUser struct {
	ID        int        `gorm:"primary_key;auto_increment;not_null"`
	Username  string     `gorm:"not_null;size:40"`
	Name      string     `gorm:"not_null;size:40"`
	Email     string     `gorm:"not_null;size:40"`
	Password  string     `gorm:"not_null;size:255"`
	Status    int        `gorm:"not_null;default:0"`
	RoleID    int        `gorm:"not_null;index:idx_role_id"`
	Role      CrRole     `gorm:"foreignkey:RoleID;association_foreignkey:ID"`
	CreatedAt time.Time  `gorm:"not_null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not_null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"null"`
}
