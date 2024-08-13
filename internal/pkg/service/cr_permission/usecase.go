package cr_permission

import "github.com/tegarsubkhan236/go-rbac/internal/pkg/model"

type UseCase interface {
	FetchAllPermission(limit, page int) ([]*model.CrPermission, error)
	FetchPermissionByID(id uint) (*model.CrPermission, error)
	CreatePermission(data *model.CrPermission) error
	UpdatePermission(id uint, data *model.CrPermission) error
	DeletePermission(id uint) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (c *useCase) FetchAllPermission(limit, page int) ([]*model.CrPermission, error) {
	return c.repo.FindAll()
}

func (c *useCase) FetchPermissionByID(id uint) (*model.CrPermission, error) {
	result, err := c.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return result.ToResponse(), err
}

func (c *useCase) CreatePermission(data *model.CrPermission) error {
	return c.repo.Create(data)
}

func (c *useCase) UpdatePermission(id uint, data *model.CrPermission) error {
	result, err := c.repo.FindByID(id)
	if err != nil {
		return err
	}

	if result.Name != "" {
		data.Name = result.Name
	}

	return c.repo.Update(data)
}

func (c *useCase) DeletePermission(id uint) error {
	return c.repo.Delete(&model.CrPermission{ID: id})
}
