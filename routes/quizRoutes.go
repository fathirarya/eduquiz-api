package routes

import (
	"eduquiz-api/controller"
	"eduquiz-api/repository"
	"eduquiz-api/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func QuizRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	quizRepository := repository.NewQuizRepository(db)
	quizService := service.NewQuizService(quizRepository, validate)
	quizController := controller.NewQuizController(quizService)

	quizGroup := e.Group("api/v1/quiz")

	quizGroup.POST("", quizController.CreateQuizController)
	quizGroup.GET("/:id", quizController.GetQuizByIdController)
	quizGroup.GET("", quizController.GetAllQuizController)
	quizGroup.GET("/", quizController.GetQuizByTitleController)
	quizGroup.PUT("/:id", quizController.UpdateQuizController)
	quizGroup.DELETE("/:id", quizController.DeleteQuizController)

}
