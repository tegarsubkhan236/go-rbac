package cr_permission

import (
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]*model.CrPermission, error)
	FindByID(id uint) (*model.CrPermission, error)
	Create(data *model.CrPermission) error
	Update(data *model.CrPermission) error
	Delete(data *model.CrPermission) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (c repository) FindAll() ([]*model.CrPermission, error) {
	var data []*model.CrPermission
	result := c.db.Find(&data)
	return data, result.Error
}

func (c repository) FindByID(id uint) (*model.CrPermission, error) {
	var item model.CrPermission
	result := c.db.First(&item, id)
	return &item, result.Error
}

func (c repository) Create(data *model.CrPermission) error {
	return c.db.Create(data).Error
}

func (c repository) Update(data *model.CrPermission) error {
	return c.db.Save(&data).Error
}

func (c repository) Delete(data *model.CrPermission) error {
	return c.db.Delete(&data).Error
}
