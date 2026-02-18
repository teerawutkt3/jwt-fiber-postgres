package routes // เปลี่ยนจาก internal เป็น routes ให้ตรงกับชื่อ folder มาตรฐาน

import (
	"fiber-poc-api/database/repository"
	"fiber-poc-api/handlers"
	"fiber-poc-api/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app fiber.Router, jwtMiddleware fiber.Handler, db *gorm.DB) {

	// ==> Repositories
	userRepo := repository.NewUserRepository(db)
	loginHistoryRepo := repository.NewLoginHistoryRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	privilegeRepo := repository.NewPrivilegeRepository(db)
	rolePrivilegeRepo := repository.NewRolePrivilegeRepository(db)

	// ==> Services
	authService := services.NewAuthService(userRepo, loginHistoryRepo)
	roleService := services.NewRoleService(userRepo, roleRepo, privilegeRepo, rolePrivilegeRepo)

	// ==> Handlers
	authHandler := handlers.NewAuthHandler(authService)
	roleHandler := handlers.NewRoleHandler(roleService)

	api := app.Group("/api/v1")

	// ==> Public Routes
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.LoginHandler)
	auth.Post("/register", authHandler.RegisterHandler)

	protected := api.Group("/", jwtMiddleware)

	// ==> User Routes
	user := protected.Group("/user")
	user.Get("/get/all", authHandler.GetUserAllHandler)

	// ==> Role Routes
	role := protected.Group("/role")
	role.Get("/create-role", roleHandler.CreateRoleHandler) // หมายเหตุ: ปกติ create ควรเป็น POST แต่ถ้า logic คุณเป็น GET ก็คงไว้ตามเดิม
	role.Post("/create-role-privilege", roleHandler.CreateRolePrivilegeHandler)

	// ==> for initial data
	roleSupport := api.Group("/role")
	roleSupport.Get("/initial-permission", roleHandler.InitialHandler)
}
