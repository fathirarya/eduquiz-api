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

func QuizRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	quizRepository := repository.NewQuizRepository(db)
	quizService := service.NewQuizService(quizRepository, validate)
	quizController := controller.NewQuizController(quizService)

	quizGroup := e.Group("api/v1/quiz")

	quizGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	quizGroup.POST("", quizController.CreateQuizController, middleware.AuthMiddleware("Guru"))
	quizGroup.GET("/:id", quizController.GetQuizByIdController)
	quizGroup.GET("", quizController.GetAllQuizController)
	quizGroup.GET("/", quizController.GetQuizByTitleController, middleware.AuthMiddleware("Guru"))
	quizGroup.PUT("/:id", quizController.UpdateQuizController, middleware.AuthMiddleware("Guru"))
	quizGroup.DELETE("/:id", quizController.DeleteQuizController, middleware.AuthMiddleware("Guru"))

}
