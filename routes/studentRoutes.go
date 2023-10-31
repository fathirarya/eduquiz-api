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

func StudentRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	studentRepository := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepository, validate)
	studentController := controller.NewStudentController(studentService)

	studentsGroup := e.Group("api/v1/students")
	studentsGroup.POST("/register", studentController.RegisterStudentController)
	studentsGroup.POST("/login", studentController.LoginStudentController)

	studentsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	studentsGroup.GET("/:id", studentController.GetStudentController, middleware.AuthMiddleware("Student"))
	studentsGroup.GET("", studentController.GetStudentsController, middleware.AuthMiddleware("Teacher"))
	studentsGroup.PUT("/:id", studentController.UpdateStudentController, middleware.AuthMiddleware("Student"))
	studentsGroup.DELETE("/:id", studentController.DeleteStudentController, middleware.AuthMiddleware("Student"))

}
