package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/go-rbac/internal/api/constant"
	"github.com/tegarsubkhan236/go-rbac/internal/api/handler"
)

type RouteConfig struct {
	App                 *fiber.App
	CrPermissionHandler handler.CrPermissionHandler
	CrRoleHandler       handler.CrRoleHandler
	CrUserHandler       handler.CrUserHandler
	CrAuthHandler       handler.CrAuthHandler
	Protected           fiber.Handler
	Gateway             func(permission ...string) fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	guest := c.App.Group("/api")
	guest.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("HALLO WORLD")
	})
	guest.Post("/register", c.CrAuthHandler.Register)
	guest.Post("/login", c.CrAuthHandler.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	api := c.App.Group("/api")
	api.Use(c.Protected)

	permission := api.Group("/permission")
	permission.Get("/", c.Gateway(constant.READ_PERMISSION), c.CrPermissionHandler.HandleFetchAllPermission)
	permission.Get("/:id", c.Gateway(constant.READ_PERMISSION), c.CrPermissionHandler.HandleFetchByIdPermission)
	permission.Post("/", c.Gateway(constant.CREATE_PERMISSION), c.CrPermissionHandler.HandleInsertPermission)
	permission.Put("/:id", c.Gateway(constant.UPDATE_PERMISSION), c.CrPermissionHandler.HandleUpdatePermission)
	permission.Delete("/:id", c.Gateway(constant.DELETE_PERMISSION), c.CrPermissionHandler.HandleDeletePermission)

	role := api.Group("/role")
	role.Get("/", c.Gateway(constant.READ_ROLE), c.CrRoleHandler.HandleFetchAllRole)
	role.Get("/:id", c.Gateway(constant.READ_ROLE), c.CrRoleHandler.HandleFetchByIdRole)
	role.Post("/", c.Gateway(constant.CREATE_ROLE), c.CrRoleHandler.HandleInsertRole)
	role.Put("/:id", c.Gateway(constant.UPDATE_ROLE), c.CrRoleHandler.HandleUpdateRole)
	role.Delete("/:id", c.Gateway(constant.DELETE_ROLE), c.CrRoleHandler.HandleDeleteRole)

	user := api.Group("/user")
	user.Get("/", c.Gateway(constant.READ_USER), c.CrUserHandler.HandleFetchAllUser)
	user.Get("/:id", c.Gateway(constant.READ_USER), c.CrUserHandler.HandleFetchByIdUser)
	user.Post("/", c.Gateway(constant.CREATE_USER), c.CrUserHandler.HandleInsertUser)
	user.Put("/:id", c.Gateway(constant.UPDATE_USER), c.CrUserHandler.HandleUpdateUser)
	user.Delete("/:id", c.Gateway(constant.DELETE_USER), c.CrUserHandler.HandleDeleteUser)
}
