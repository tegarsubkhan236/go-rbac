package cr_permission

import "github.com/tegarsubkhan236/redis-jwt-auth/internal/pkg/model"

type CrPermissionUseCase interface {
	FetchAllPermission() ([]*model.CrPermission, error)
	FetchPermissionByID(id uint) (*model.CrPermission, error)
	CreatePermission(data *model.CrPermission) error
	UpdatePermission(id uint, data *model.CrPermission) error
	DeletePermission(id uint) error
}

type crPermissionUseCase struct {
	repo CrPermissionRepository
}

func NewCrPermissionUseCase(repo CrPermissionRepository) CrPermissionUseCase {
	return &crPermissionUseCase{
		repo: repo,
	}
}

func (c *crPermissionUseCase) FetchAllPermission() ([]*model.CrPermission, error) {
	return c.repo.FindAll()
}

func (c *crPermissionUseCase) FetchPermissionByID(id uint) (*model.CrPermission, error) {
	result, err := c.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return result.ToResponse(), err
}

func (c *crPermissionUseCase) CreatePermission(data *model.CrPermission) error {
	return c.repo.Create(data)
}

func (c *crPermissionUseCase) UpdatePermission(id uint, data *model.CrPermission) error {
	result, err := c.repo.FindByID(id)
	if err != nil {
		return err
	}

	if result.Name != "" {
		data.Name = result.Name
	}

	return c.repo.Update(data)
}

func (c *crPermissionUseCase) DeletePermission(id uint) error {
	return c.repo.Delete(&model.CrPermission{ID: id})
}
