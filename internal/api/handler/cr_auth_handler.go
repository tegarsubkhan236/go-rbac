package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/service/cr_user"
)

type CrAuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type crAuthHandler struct {
	usecase cr_user.UseCase
}

func NewCrAuthHandler(usecase cr_user.UseCase) CrAuthHandler {
	return &crAuthHandler{
		usecase: usecase,
	}
}

func (r *crAuthHandler) Login(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crAuthHandler) Register(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (r *crAuthHandler) Logout(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
