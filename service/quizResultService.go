package service

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/web"
	"eduquiz-api/repository"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type QuizResultService interface {
	PostResult(ctx echo.Context, request web.QuizResultCreateRequest) (*domain.QuizResult, error)
}

type QuizResultServiceImpl struct {
	QuizResultRepository   repository.QuizResultRepository
	AttemptAnswerRepostory repository.AttemptAnswerRepository
	Validate               *validator.Validate
}

func NewQuizResultService(quizResultRepository repository.QuizResultRepository, attemptRepository repository.AttemptAnswerRepository, validate *validator.Validate) *QuizResultServiceImpl {
	return &QuizResultServiceImpl{
		QuizResultRepository:   quizResultRepository,
		AttemptAnswerRepostory: attemptRepository,
		Validate:               validate,
	}
}

// func (service *QuizResultServiceImpl) PostResult(ctx echo.Context, request web.QuizResultCreateRequest) (*domain.QuizResult, error) {
// }
