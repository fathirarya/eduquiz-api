package routes

// func QuizResultRoutes(e *echo.Echo, DB *gorm.DB, validate *validator.Validate) {
// 	attemptAnswerRepository := repository.NewAttemptAnswerRepository(DB)
// 	keyAnswerRepository := repository.NewKeyAnswerRepository(DB)
// 	attemptAnswerService := service.NewAttemptAnswerService(attemptAnswerRepository, keyAnswerRepository, validate)

// 	quizResultRepository := repository.NewQuizResultRepository(DB)
// 	quizResultService := service.NewQuizResultService(quizResultRepository, attemptAnswerRepository, validate)
// 	// quizResultController := controller.NewQuizResultController(quizResultService, attemptAnswerService)

// 	quizResult := e.Group("api/v1/quizresult")

// 	quizResult.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

// 	quizResult.POST("", quizResultController.CreateQuizResultController)
// }
