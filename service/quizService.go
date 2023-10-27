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

type QuizService interface {
	CreateQuiz(ctx echo.Context, request web.QuizCreateRequest) (*domain.Quiz, error)
	UpdateQuiz(ctx echo.Context, request web.QuizUpdateRequest, id int) (*domain.Quiz, error)
	FindQuizByTitle(ctx echo.Context, Title string) (*domain.Quiz, error)
	FindQuizById(ctx echo.Context, id int) (*domain.Quiz, error)
	FindAllQuiz(ctx echo.Context) ([]domain.Quiz, error)
	DeleteQuiz(ctx echo.Context, id int) error
}

type QuizServiceImpl struct {
	QuizRepository repository.QuizRepository
	Validate       *validator.Validate
}

func NewQuizService(quizRepository repository.QuizRepository, validate *validator.Validate) *QuizServiceImpl {
	return &QuizServiceImpl{
		QuizRepository: quizRepository,
		Validate:       validate,
	}
}

func (service *QuizServiceImpl) CreateQuiz(ctx echo.Context, request web.QuizCreateRequest) (*domain.Quiz, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	existingQuiz, _ := service.QuizRepository.FindByTitle(request.Title)
	if existingQuiz != nil {
		return nil, fmt.Errorf("Quiz Already Exist")
	}
	quiz := req.QuizCreateRequestToQuizDomain(request)
	result, err := service.QuizRepository.Create(quiz)
	if err != nil {
		return nil, fmt.Errorf("Error when creating Quiz: %s", err.Error())
	}

	return result, nil
}

func (service *QuizServiceImpl) UpdateQuiz(ctx echo.Context, request web.QuizUpdateRequest, id int) (*domain.Quiz, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingQuiz, _ := service.QuizRepository.FindById(id)
	if existingQuiz == nil {
		return nil, fmt.Errorf("Quiz Not Found")
	}

	quiz := req.QuizUpdateRequestToQuizDomain(request)
	result, err := service.QuizRepository.Update(quiz, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating Quiz: %s", err.Error())
	}

	return result, nil
}

func (service *QuizServiceImpl) FindQuizByTitle(ctx echo.Context, Title string) (*domain.Quiz, error) {

	existingQuiz, _ := service.QuizRepository.FindByTitle(Title)
	if existingQuiz == nil {
		return nil, fmt.Errorf("Quiz Not Found")
	}

	return existingQuiz, nil
}

func (service *QuizServiceImpl) FindQuizById(ctx echo.Context, id int) (*domain.Quiz, error) {

	existingQuiz, _ := service.QuizRepository.FindById(id)
	if existingQuiz == nil {
		return nil, fmt.Errorf("Quiz Not Found")
	}

	return existingQuiz, nil
}

func (service *QuizServiceImpl) FindAllQuiz(ctx echo.Context) ([]domain.Quiz, error) {

	existingQuiz, _ := service.QuizRepository.FindAll()
	if existingQuiz == nil {
		return nil, fmt.Errorf("Quizs Not Found")
	}

	return existingQuiz, nil
}

func (service *QuizServiceImpl) DeleteQuiz(ctx echo.Context, id int) error {

	existingQuiz, _ := service.QuizRepository.FindById(id)
	if existingQuiz == nil {
		return fmt.Errorf("Quiz Not Found")
	}

	err := service.QuizRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting Quiz: %s", err)
	}

	return nil
}
