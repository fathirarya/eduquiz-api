package routes

import (
	"eduquiz-api/controller"
	"eduquiz-api/repository"
	"eduquiz-api/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func QuizResultRoutes(e *echo.Echo, DB *gorm.DB, validate *validator.Validate) {
	attemptAnswerRepository := repository.NewAttemptAnswerRepository(DB)
	keyAnswerRepository := repository.NewKeyAnswerRepository(DB)
	questionRepository := repository.NewQuestionRepository(DB)
	attemptAnswerService := service.NewAttemptAnswerService(attemptAnswerRepository, keyAnswerRepository, validate)

	quizResultRepository := repository.NewQuizResultRepository(DB)
	quizResultService := service.NewQuizResultService(quizResultRepository, attemptAnswerRepository, questionRepository, validate)
	quizResultController := controller.NewQuizResultController(quizResultService, attemptAnswerService)

	quizResult := e.Group("api/v1/quizresult")

	quizResult.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	quizResult.POST("", quizResultController.CreateQuizResultController)
}
