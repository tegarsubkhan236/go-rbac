package cr_user

import "github.com/tegarsubkhan236/go-rbac/internal/pkg/model"

type UseCase interface {
	FetchAllUser(limit, page int) ([]*model.CrUser, error)
	FetchUserByID(id uint) (*model.CrUser, error)
	CreateUser(data *model.CrUser) error
	UpdateUser(id uint, data *model.CrUser) error
	DeleteUser(id uint) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (u *useCase) FetchAllUser(limit, page int) ([]*model.CrUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) FetchUserByID(id uint) (*model.CrUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) CreateUser(data *model.CrUser) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) UpdateUser(id uint, data *model.CrUser) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) DeleteUser(id uint) error {
	//TODO implement me
	panic("implement me")
}
