package main

import (
	"eduquiz-api/config"
	"eduquiz-api/routes"
	"net/http"
        "log"
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


	DB := config.ConnectDB()

	myApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to EduQuiz API Services")
	})

	routes.StudentRoutes(myApp, DB, validate)
	routes.TeacherRoutes(myApp, DB, validate)
	routes.QuizCategoriesRoutes(myApp, DB, validate)
	routes.QuizRoutes(myApp, DB, validate)
	routes.QuestionRoutes(myApp, DB, validate)
	routes.KeyAnswerRoutes(myApp, DB, validate)
	routes.AttemptAnswerRoutes(myApp, DB, validate)
	routes.QuizResultRoutes(myApp, DB, validate)

	myApp.Pre(middleware.RemoveTrailingSlash())
	myApp.Use(middleware.CORS())
	myApp.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	myApp.Logger.Fatal(myApp.Start(":8080"))
}
