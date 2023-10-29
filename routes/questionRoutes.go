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

func QuestionRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	questionRepository := repository.NewQuestionRepository(db)
	questionService := service.NewQuestionService(questionRepository, validate)
	questionController := controller.NewQuestionController(questionService)

	questionGroup := e.Group("api/v1/question")
	questionGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	questionGroup.POST("", questionController.CreateQuestionController, middleware.AuthMiddleware("Guru"))
	questionGroup.GET("/:id", questionController.GetQuestionByIdController)
	questionGroup.GET("", questionController.GetAllQuestionController)
	questionGroup.PUT("/:id", questionController.UpdateQuestionController, middleware.AuthMiddleware("Guru"))
	questionGroup.DELETE("/:id", questionController.DeleteQuestionController, middleware.AuthMiddleware("Guru"))

}
