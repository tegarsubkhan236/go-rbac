package cr_role

import "github.com/tegarsubkhan236/go-rbac/internal/pkg/model"

type UseCase interface {
	FetchAllRole(limit, page int) ([]*model.CrRole, error)
	FetchRoleByID(id uint) (*model.CrRole, error)
	CreateRole(data *model.CrRole) error
	UpdateRole(id uint, data *model.CrRole) error
	DeleteRole(id uint) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (u *useCase) FetchAllRole(limit, page int) ([]*model.CrRole, error) {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) FetchRoleByID(id uint) (*model.CrRole, error) {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) CreateRole(data *model.CrRole) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) UpdateRole(id uint, data *model.CrRole) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) DeleteRole(id uint) error {
	//TODO implement me
	panic("implement me")
}
