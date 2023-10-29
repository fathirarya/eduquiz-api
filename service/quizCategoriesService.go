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

type QuizCategoryService interface {
	CreateQuizCategory(ctx echo.Context, request web.QuizCategoryCreateRequest) (*domain.QuizCategory, error)
	FindQuizCategoryById(ctx echo.Context, id int) (*domain.QuizCategory, error)
	FindAllQuizCategory(ctx echo.Context) ([]domain.QuizCategory, error)
	DeleteQuizCategory(ctx echo.Context, id int) error
}

type QuizCategoryServiceImpl struct {
	QuizCategoryRepository repository.QuizCategoryRepository
	Validate               *validator.Validate
}

func NewQuizCategoryService(QuizCategoryRepository repository.QuizCategoryRepository, validate *validator.Validate) *QuizCategoryServiceImpl {
	return &QuizCategoryServiceImpl{
		QuizCategoryRepository: QuizCategoryRepository,
		Validate:               validate,
	}
}

func (service *QuizCategoryServiceImpl) CreateQuizCategory(ctx echo.Context, request web.QuizCategoryCreateRequest) (*domain.QuizCategory, error) {
	err := service.Validate.Struct(request)

	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	existingQuizCategory, _ := service.QuizCategoryRepository.FindByName(request.Category)
	if existingQuizCategory != nil {
		return nil, fmt.Errorf("Category Already Exist")
	}
	fmt.Println(request)
	quizCategory := req.QuizCategoryCreateRequestToQuizCategoryDomain(request)
	fmt.Println(quizCategory)
	result, err := service.QuizCategoryRepository.Create(quizCategory)
	if err != nil {
		return nil, fmt.Errorf("Error when creating QuizCategory: %s", err.Error())
	}

	return result, nil
}

func (service *QuizCategoryServiceImpl) FindQuizCategoryById(ctx echo.Context, id int) (*domain.QuizCategory, error) {

	existingUser, _ := service.QuizCategoryRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("Category Not Found")
	}

	return existingUser, nil
}

func (service *QuizCategoryServiceImpl) FindAllQuizCategory(ctx echo.Context) ([]domain.QuizCategory, error) {

	existingUser, _ := service.QuizCategoryRepository.FindAll()
	if existingUser == nil {
		return nil, fmt.Errorf("Categories Not Found")
	}

	return existingUser, nil
}

func (service *QuizCategoryServiceImpl) DeleteQuizCategory(ctx echo.Context, id int) error {

	existingUser, _ := service.QuizCategoryRepository.FindById(id)
	if existingUser == nil {
		return fmt.Errorf("Category Not Found")
	}

	err := service.QuizCategoryRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting Category: %s", err.Error())
	}

	return nil
}
