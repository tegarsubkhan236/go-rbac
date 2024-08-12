package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/redis-jwt-auth/internal/pkg/service/cr_permission"
	"github.com/tegarsubkhan236/redis-jwt-auth/internal/util"
)

func HandleFetchAllPermission(usecase cr_permission.CrPermissionUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := usecase.FetchAllPermission()
		if err != nil {
			return util.ResponseInternalServerError(c, err.Error())
		}

		return util.ResponseOKWithPages(c, 0, 0, 0, result)
	}
}

func HandleFetchByIdPermission(usecase cr_permission.CrPermissionUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := usecase.FetchPermissionByID(1)
		if err != nil {
			return util.ResponseInternalServerError(c, err.Error())
		}

		return util.ResponseOK(c, result)
	}
}

func HandleInsertPermission(usecase cr_permission.CrPermissionUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := usecase.FetchPermissionByID(1)
		if err != nil {
			return util.ResponseInternalServerError(c, err.Error())
		}

		return util.ResponseCreated(c, result)
	}
}

func HandleUpdatePermission(usecase cr_permission.CrPermissionUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := usecase.FetchPermissionByID(1)
		if err != nil {
			return util.ResponseInternalServerError(c, err.Error())
		}

		return util.ResponseOK(c, result)
	}
}

func HandleDeletePermission(usecase cr_permission.CrPermissionUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := usecase.FetchPermissionByID(1)
		if err != nil {
			return util.ResponseInternalServerError(c, err.Error())
		}

		return util.ResponseOK(c, result)
	}
}
