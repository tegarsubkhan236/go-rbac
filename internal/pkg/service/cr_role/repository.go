package cr_role

import (
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]*model.CrRole, error)
	FindByID(id uint) (*model.CrRole, error)
	Create(data *model.CrRole) error
	Update(data *model.CrRole) error
	Delete(data *model.CrRole) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll() ([]*model.CrRole, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindByID(id uint) (*model.CrRole, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Create(data *model.CrRole) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(data *model.CrRole) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(data *model.CrRole) error {
	//TODO implement me
	panic("implement me")
}
