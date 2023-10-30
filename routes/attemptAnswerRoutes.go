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

func AttemptAnswerRoutes(e *echo.Echo, DB *gorm.DB, validate *validator.Validate) {
	keyAnswerRepository := repository.NewKeyAnswerRepository(DB)

	attemptAnswerRepository := repository.NewAttemptAnswerRepository(DB)
	attemptAnswerService := service.NewAttemptAnswerService(attemptAnswerRepository, keyAnswerRepository, validate)
	attemptAnswerController := controller.NewAttemptAnswerController(attemptAnswerService)

	attemptAnswer := e.Group("api/v1/attempt-answer")

	attemptAnswer.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	attemptAnswer.POST("", attemptAnswerController.CreateAttemptAnswerController, middleware.AuthMiddleware("Student"))
	attemptAnswer.GET("/:id", attemptAnswerController.GetAttemptAnswerByIdController, middleware.AuthMiddleware("Student"))
	attemptAnswer.GET("", attemptAnswerController.GetAllAttemptAnswerController, middleware.AuthMiddleware("Student"))
}
