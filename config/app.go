package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis/v3"
	"github.com/spf13/viper"
	"github.com/tegarsubkhan236/go-rbac/internal/api/handler"
	"github.com/tegarsubkhan236/go-rbac/internal/api/middleware"
	"github.com/tegarsubkhan236/go-rbac/internal/api/route"
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/service/cr_permission"
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/service/cr_role"
	"github.com/tegarsubkhan236/go-rbac/internal/pkg/service/cr_user"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App       *fiber.App
	DB        *gorm.DB
	Config    *viper.Viper
	Redis     *redis.Storage
	Validator *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	// crPermission
	crPermissionRepository := cr_permission.NewRepository(config.DB)
	crPermissionUseCase := cr_permission.NewUseCase(crPermissionRepository)
	crPermissionHandler := handler.NewCrPermissionHandler(crPermissionUseCase)

	// crRole
	crRoleRepository := cr_role.NewRepository(config.DB)
	crRoleUseCase := cr_role.NewUseCase(crRoleRepository)
	crRoleHandler := handler.NewCrRoleHandler(crRoleUseCase)

	// crUser
	crUserRepository := cr_user.NewRepository(config.DB)
	crUserUseCase := cr_user.NewUseCase(crUserRepository)
	crUserHandler := handler.NewCrUserHandler(crUserUseCase)

	// crAuth
	crAuthHandler := handler.NewCrAuthHandler(crUserUseCase)

	// Setup Middleware
	protectedMiddleware := middleware.Protected(config.Config)
	gatewayMiddleware := middleware.Gateway

	// Setup Route
	routeConfig := route.RouteConfig{
		App:                 config.App,
		CrPermissionHandler: crPermissionHandler,
		CrRoleHandler:       crRoleHandler,
		CrUserHandler:       crUserHandler,
		CrAuthHandler:       crAuthHandler,
		Protected:           protectedMiddleware,
		Gateway:             gatewayMiddleware,
	}
	routeConfig.Setup()
}
