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
	QuestionRepository     repository.QuestionRepository
	Validate               *validator.Validate
}

func NewQuizResultService(quizResultRepository repository.QuizResultRepository, attemptRepository repository.AttemptAnswerRepository, questionRepository repository.QuestionRepository, validate *validator.Validate) *QuizResultServiceImpl {
	return &QuizResultServiceImpl{
		QuizResultRepository:   quizResultRepository,
		AttemptAnswerRepostory: attemptRepository,
		QuestionRepository:     questionRepository,
		Validate:               validate,
	}
}

func (service *QuizResultServiceImpl) PostResult(ctx echo.Context, request web.QuizResultCreateRequest) (*domain.QuizResult, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	score, err := service.calculateScore(request)
	if err != nil {
		return nil, err
	}

	quizResult := domain.QuizResult{
		StudentID: request.StudentID,
		QuizID:    request.QuizID,
		Score:     int(score),
	}

	return service.QuizResultRepository.PostResult(&quizResult)
}

func (service *QuizResultServiceImpl) calculateScore(request web.QuizResultCreateRequest) (float64, error) {
	attemptAnswers, err := service.AttemptAnswerRepostory.FindByStudentId(int(request.StudentID))
	if err != nil {
		return 0, err
	}

	var correctAnswers int

	for _, attemptAnswer := range attemptAnswers {
		if attemptAnswer.IsCorrect {
			correctAnswers++
		}
	}

	totalQuestions, err := service.QuestionRepository.FindByQuizId(int(request.QuizID))

	score := float64(correctAnswers) / float64(len(totalQuestions)) * 100

	return score, nil
}
