package main

import (
	"eduquiz-api/config"
	"eduquiz-api/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	myApp := echo.New()
	validate := validator.New()

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}
	config.ConnectDB()
	config.Migrate()

	myApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to EduQuiz API Services")
	})

	routes.StudentRoutes(myApp, config.DB, validate)
	routes.TeacherRoutes(myApp, config.DB, validate)
	routes.QuizCategoriesRoutes(myApp, config.DB, validate)
	routes.QuizRoutes(myApp, config.DB, validate)
	routes.QuestionRoutes(myApp, config.DB, validate)
	routes.KeyAnswerRoutes(myApp, config.DB, validate)
	routes.AttemptAnswerRoutes(myApp, config.DB, validate)
	routes.QuizResultRoutes(myApp, config.DB, validate)

	myApp.Pre(middleware.RemoveTrailingSlash())
	myApp.Use(middleware.CORS())
	myApp.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	myApp.Logger.Fatal(myApp.Start(":8080"))
}
