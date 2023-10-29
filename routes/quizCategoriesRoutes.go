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

func QuizCategoriesRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	quizCategoryRepository := repository.NewQuizCategoryRepository(db)
	quizCategoryService := service.NewQuizCategoryService(quizCategoryRepository, validate)
	quizCategoryController := controller.NewQuizCategoryController(quizCategoryService)

	quizCategoryGroup := e.Group("api/v1/quiz-category")

	quizCategoryGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	quizCategoryGroup.POST("", quizCategoryController.CreateQuizCategoryController, middleware.AuthMiddleware("Guru"))
	quizCategoryGroup.GET("/:id", quizCategoryController.GetQuizCategoryByIdController, middleware.AuthMiddleware("Guru"))
	quizCategoryGroup.GET("", quizCategoryController.GetAllQuizCategoryController, middleware.AuthMiddleware("Guru"))
	quizCategoryGroup.DELETE("/:id", quizCategoryController.DeleteQuizCategoryController, middleware.AuthMiddleware("Guru"))

}
