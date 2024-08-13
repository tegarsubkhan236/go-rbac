package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/service/cr_user"
)

type CrUserHandler interface {
	HandleFetchAllUser(c *fiber.Ctx) error
	HandleFetchByIdUser(c *fiber.Ctx) error
	HandleInsertUser(c *fiber.Ctx) error
	HandleUpdateUser(c *fiber.Ctx) error
	HandleDeleteUser(c *fiber.Ctx) error
}

type crUserHandler struct {
	usecase cr_user.UseCase
}

func NewCrUserHandler(usecase cr_user.UseCase) CrUserHandler {
	return &crUserHandler{
		usecase: usecase,
	}
}

func (r *crUserHandler) HandleFetchAllUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crUserHandler) HandleFetchByIdUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crUserHandler) HandleInsertUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crUserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crUserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
