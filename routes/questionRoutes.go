package routes

import (
	"eduquiz-api/controller"
	"eduquiz-api/repository"
	"eduquiz-api/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func QuestionRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	questionRepository := repository.NewQuestionRepository(db)
	questionService := service.NewQuestionService(questionRepository, validate)
	questionController := controller.NewQuestionController(questionService)

	questionGroup := e.Group("api/v1/question")

	questionGroup.POST("", questionController.CreateQuestionController)
	questionGroup.GET("/:id", questionController.GetQuestionByIdController)
	questionGroup.GET("", questionController.GetAllQuestionController)
	questionGroup.PUT("/:id", questionController.UpdateQuestionController)
	questionGroup.DELETE("/:id", questionController.DeleteQuestionController)

}
