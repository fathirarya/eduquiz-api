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

type QuestionService interface {
	CreateQuestion(ctx echo.Context, request web.QuestionCreateRequest) (*domain.Question, error)
	UpdateQuestion(ctx echo.Context, request web.QuestionUpdateRequest, id int) (*domain.Question, error)
	FindByIdQuestion(ctx echo.Context, id int) (*domain.Question, error)
	FindAllQuestion(ctx echo.Context) ([]domain.Question, error)
	DeleteQuestion(ctx echo.Context, id int) error
}

type QuestionServiceImpl struct {
	QuestionRepository repository.QuestionRepository
	Validate           *validator.Validate
}

func NewQuestionService(QuestionRepository repository.QuestionRepository, validate *validator.Validate) *QuestionServiceImpl {
	return &QuestionServiceImpl{
		QuestionRepository: QuestionRepository,
		Validate:           validate,
	}
}

func (service *QuestionServiceImpl) CreateQuestion(ctx echo.Context, request web.QuestionCreateRequest) (*domain.Question, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingQuestion, _ := service.QuestionRepository.FindByQuestion(request.Question)
	if existingQuestion != nil {
		return nil, fmt.Errorf("Question Already Exist")
	}
	question := req.QuestionCreateRequestToQuestionDomain(request)

	result, err := service.QuestionRepository.Create(question)
	if err != nil {
		return nil, fmt.Errorf("Error when creating Question: %s", err.Error())
	}

	return result, nil
}

func (service *QuestionServiceImpl) UpdateQuestion(ctx echo.Context, request web.QuestionUpdateRequest, id int) (*domain.Question, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingQuestion, _ := service.QuestionRepository.FindById(id)
	if existingQuestion == nil {
		return nil, fmt.Errorf("Question with id %d not found", id)
	}

	question := req.QuestionUpdateRequestToQuestionDomain(request)

	result, err := service.QuestionRepository.Update(question, id)
	if err != nil {
		return nil, fmt.Errorf("Error when update question: %w", err)
	}

	return result, nil
}

func (service *QuestionServiceImpl) FindByIdQuestion(ctx echo.Context, id int) (*domain.Question, error) {

	question, _ := service.QuestionRepository.FindById(id)
	if question == nil {
		return nil, fmt.Errorf("Question with id %d not found", id)
	}

	return question, nil
}

func (service *QuestionServiceImpl) FindAllQuestion(ctx echo.Context) ([]domain.Question, error) {
	question, err := service.QuestionRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Error when find all question: %w", err)
	}

	return question, nil
}

func (service *QuestionServiceImpl) DeleteQuestion(ctx echo.Context, id int) error {
	question, _ := service.QuestionRepository.FindById(id)
	if question != nil {
		return fmt.Errorf("Question with id %d not found", id)
	}

	err := service.QuestionRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when delete question: %w", err)
	}

	return nil
}
