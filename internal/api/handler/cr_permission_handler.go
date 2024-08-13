package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/service/cr_permission"
)

type CrPermissionHandler interface {
	HandleFetchAllPermission(c *fiber.Ctx) error
	HandleFetchByIdPermission(c *fiber.Ctx) error
	HandleInsertPermission(c *fiber.Ctx) error
	HandleUpdatePermission(c *fiber.Ctx) error
	HandleDeletePermission(c *fiber.Ctx) error
}

type crPermissionHandler struct {
	usecase cr_permission.UseCase
}

func NewCrPermissionHandler(usecase cr_permission.UseCase) CrPermissionHandler {
	return &crPermissionHandler{
		usecase: usecase,
	}
}

func (r *crPermissionHandler) HandleFetchAllPermission(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	page := c.QueryInt("page")

	result, err := r.usecase.FetchAllPermission(limit, page)
	if err != nil {
		return ResponseInternalServerError(c, err.Error())
	}

	return ResponseOKWithPages(c, 0, 0, 0, result)
}

func (r *crPermissionHandler) HandleFetchByIdPermission(c *fiber.Ctx) error {
	result, err := r.usecase.FetchPermissionByID(1)
	if err != nil {
		return ResponseInternalServerError(c, err.Error())
	}

	return ResponseOK(c, result)
}

func (r *crPermissionHandler) HandleInsertPermission(c *fiber.Ctx) error {
	result, err := r.usecase.FetchPermissionByID(1)
	if err != nil {
		return ResponseInternalServerError(c, err.Error())
	}

	return ResponseCreated(c, result)
}

func (r *crPermissionHandler) HandleUpdatePermission(c *fiber.Ctx) error {
	result, err := r.usecase.FetchPermissionByID(1)
	if err != nil {
		return ResponseInternalServerError(c, err.Error())
	}

	return ResponseOK(c, result)
}

func (r *crPermissionHandler) HandleDeletePermission(c *fiber.Ctx) error {
	result, err := r.usecase.FetchPermissionByID(1)
	if err != nil {
		return ResponseInternalServerError(c, err.Error())
	}

	return ResponseOK(c, result)
}
