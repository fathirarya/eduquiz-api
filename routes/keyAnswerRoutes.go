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

func KeyAnswerRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	keyAnswerRepository := repository.NewKeyAnswerRepository(db)
	keyAnswerService := service.NewKeyAnswerService(keyAnswerRepository, validate)
	keyAnswerController := controller.NewKeyAnswerController(keyAnswerService)

	keyAnswerGroup := e.Group("api/v1/key-answer")

	keyAnswerGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	keyAnswerGroup.POST("", keyAnswerController.CreateKeyAnswerController, middleware.AuthMiddleware("Guru"))
	keyAnswerGroup.GET("/:id", keyAnswerController.GetKeyAnswerByIdController, middleware.AuthMiddleware("Guru"))
	keyAnswerGroup.GET("", keyAnswerController.GetAllKeyAnswerController, middleware.AuthMiddleware("Guru"))
	keyAnswerGroup.PUT("/:id", keyAnswerController.UpdateKeyAnswerController, middleware.AuthMiddleware("Guru"))
	keyAnswerGroup.DELETE("/:id", keyAnswerController.DeleteKeyAnswerController, middleware.AuthMiddleware("Guru"))

}
