package service

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/web"
	"eduquiz-api/repository"
	"eduquiz-api/utils/helper"
	"eduquiz-api/utils/req"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AttemptAnswerService interface {
	CreateAttemptAnswer(ctx echo.Context, request web.AttemptAnswerCreateReq) (*domain.AttemptAnswer, error)
	FindByIdAttemptAnswer(ctx echo.Context, id int) (*domain.AttemptAnswer, error)
	FindAllAttemptAnswer(ctx echo.Context) ([]domain.AttemptAnswer, error)
}

type AttemptAnswerServiceImpl struct {
	AttemptAnswerRepository repository.AttemptAnswerRepository
	KeyAnswerRepository     repository.KeyAnswerRepository
	Validate                *validator.Validate
}

func NewAttemptAnswerService(attemptAnswerRepository repository.AttemptAnswerRepository, keyAnswer repository.KeyAnswerRepository, validate *validator.Validate) *AttemptAnswerServiceImpl {
	return &AttemptAnswerServiceImpl{
		AttemptAnswerRepository: attemptAnswerRepository,
		KeyAnswerRepository:     keyAnswer,
		Validate:                validate,
	}
}

func (service *AttemptAnswerServiceImpl) CreateAttemptAnswer(ctx echo.Context, request web.AttemptAnswerCreateReq) (*domain.AttemptAnswer, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	attemptAnswer := req.AttemptAnswerCreateRequestToAttemptAnswerDomain(request)
	correctAnswer, _ := service.KeyAnswerRepository.FindById(int(attemptAnswer.QuestionID))
	if correctAnswer.Answer == attemptAnswer.Answer {
		attemptAnswer.IsCorrect = true
	} else {
		attemptAnswer.IsCorrect = false
	}
	result, err := service.AttemptAnswerRepository.PostAnswer(attemptAnswer)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *AttemptAnswerServiceImpl) FindByIdAttemptAnswer(ctx echo.Context, id int) (*domain.AttemptAnswer, error) {

	question, _ := service.AttemptAnswerRepository.FindById(id)
	if question == nil {
		return nil, fmt.Errorf("Answer with id %d not found", id)
	}

	return question, nil
}

func (service *AttemptAnswerServiceImpl) FindAllAttemptAnswer(ctx echo.Context) ([]domain.AttemptAnswer, error) {

	questions, err := service.AttemptAnswerRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Answer with id %d not found", err)
	}

	return questions, nil
}
