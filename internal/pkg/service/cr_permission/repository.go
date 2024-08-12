package cr_permission

import (
	"github.com/tegarsubkhan236/redis-jwt-auth/internal/pkg/model"
	"gorm.io/gorm"
)

type CrPermissionRepository interface {
	FindAll() ([]*model.CrPermission, error)
	FindByID(id uint) (*model.CrPermission, error)
	Create(data *model.CrPermission) error
	Update(data *model.CrPermission) error
	Delete(data *model.CrPermission) error
}

type crPermissionRepository struct {
	db *gorm.DB
}

func NewCrPermissionRepository(db *gorm.DB) CrPermissionRepository {
	return &crPermissionRepository{
		db: db,
	}
}

func (c crPermissionRepository) FindAll() ([]*model.CrPermission, error) {
	var data []*model.CrPermission
	result := c.db.Find(&data)
	return data, result.Error
}

func (c crPermissionRepository) FindByID(id uint) (*model.CrPermission, error) {
	var item model.CrPermission
	result := c.db.First(&item, id)
	return &item, result.Error
}

func (c crPermissionRepository) Create(data *model.CrPermission) error {
	return c.db.Create(data).Error
}

func (c crPermissionRepository) Update(data *model.CrPermission) error {
	return c.db.Save(&data).Error
}

func (c crPermissionRepository) Delete(data *model.CrPermission) error {
	return c.db.Delete(&data).Error
}
