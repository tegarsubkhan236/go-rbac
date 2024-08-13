package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/service/cr_role"
)

type CrRoleHandler interface {
	HandleFetchAllRole(c *fiber.Ctx) error
	HandleFetchByIdRole(c *fiber.Ctx) error
	HandleInsertRole(c *fiber.Ctx) error
	HandleUpdateRole(c *fiber.Ctx) error
	HandleDeleteRole(c *fiber.Ctx) error
}

type crRoleHandler struct {
	usecase cr_role.UseCase
}

func NewCrRoleHandler(usecase cr_role.UseCase) CrRoleHandler {
	return &crRoleHandler{
		usecase: usecase,
	}
}

func (r *crRoleHandler) HandleFetchAllRole(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crRoleHandler) HandleFetchByIdRole(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crRoleHandler) HandleInsertRole(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crRoleHandler) HandleUpdateRole(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crRoleHandler) HandleDeleteRole(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
