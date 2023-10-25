package routes

import (
	"eduquiz-api/controller"
	"eduquiz-api/repository"
	"eduquiz-api/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func QuizCategoriesRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	quizCategoryRepository := repository.NewQuizCategoryRepository(db)
	quizCategoryService := service.NewQuizCategoryService(quizCategoryRepository, validate)
	quizCategoryController := controller.NewQuizCategoryController(quizCategoryService)

	quizCategoryGroup := e.Group("api/v1/quiz-category")

	quizCategoryGroup.POST("", quizCategoryController.CreateQuizCategoryController)
	quizCategoryGroup.GET("/:id", quizCategoryController.GetQuizCategoryByIdController)
	quizCategoryGroup.GET("", quizCategoryController.GetAllQuizCategoryController)
	quizCategoryGroup.GET("/:name", quizCategoryController.GetQuizCategoryByNameController)
	quizCategoryGroup.DELETE("/:id", quizCategoryController.DeleteQuizCategoryController)

}
