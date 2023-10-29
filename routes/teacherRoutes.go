package routes

import (
	"eduquiz-api/controller"
	"eduquiz-api/repository"
	"eduquiz-api/service"
	"eduquiz-api/utils/helper/middleware"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TeacherRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	teacherRepository := repository.NewTeacherRepository(db)
	teacherService := service.NewTeacherService(teacherRepository, validate)
	teacherController := controller.NewTeacherController(teacherService)

	teachersGroup := e.Group("api/v1/teachers")
	teachersGroup.POST("/register", teacherController.RegisterTeacherController)
	teachersGroup.POST("/login", teacherController.LoginTeacherController)

	teachersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	teachersGroup.GET("/:id", teacherController.GetTeacherController, middleware.AuthMiddleware("Guru"))
	teachersGroup.GET("", teacherController.GetTeachersController, middleware.AuthMiddleware("Guru"))
	teachersGroup.GET("/:name", teacherController.GetTeacherByNameController, middleware.AuthMiddleware("Guru"))
	teachersGroup.PUT("/:id", teacherController.UpdateTeacherController, middleware.AuthMiddleware("Guru"))
	teachersGroup.DELETE("/:id", teacherController.DeleteTeacherController, middleware.AuthMiddleware("Guru"))

}
