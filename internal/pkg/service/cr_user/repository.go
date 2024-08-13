package cr_user

import (
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]*model.CrUser, error)
	FindByID(id uint) (*model.CrUser, error)
	Create(data *model.CrUser) error
	Update(data *model.CrUser) error
	Delete(data *model.CrUser) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll() ([]*model.CrUser, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindByID(id uint) (*model.CrUser, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Create(data *model.CrUser) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(data *model.CrUser) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(data *model.CrUser) error {
	//TODO implement me
	panic("implement me")
}
