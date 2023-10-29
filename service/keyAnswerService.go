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

type KeyAnswerService interface {
	CreateAnswer(ctx echo.Context, request web.KeyCreateAnswerReq) (*domain.KeyAnswer, error)
	UpdateAnswer(ctx echo.Context, request web.KeyUpdateAnswerReq, id int) (*domain.KeyAnswer, error)
	FindAnswerById(ctx echo.Context, id int) (*domain.KeyAnswer, error)
	FindAllAnswer(ctx echo.Context) ([]domain.KeyAnswer, error)
	DeleteAnswer(ctx echo.Context, id int) error
}

type KeyAnswerServiceImpl struct {
	KeyAnswerRepository repository.KeyAnswerRepository
	Validate            *validator.Validate
}

func NewKeyAnswerService(keyAnswerRepository repository.KeyAnswerRepository, validate *validator.Validate) *KeyAnswerServiceImpl {
	return &KeyAnswerServiceImpl{
		KeyAnswerRepository: keyAnswerRepository,
		Validate:            validate,
	}
}

func (service *KeyAnswerServiceImpl) CreateAnswer(ctx echo.Context, request web.KeyCreateAnswerReq) (*domain.KeyAnswer, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	keyAnswer := req.KeyCreateAnswerRequestToKeyAnswerDomain(request)
	result, err := service.KeyAnswerRepository.Create(keyAnswer)
	if err != nil {
		return nil, fmt.Errorf("Error when creating Answer: %s", err.Error())
	}

	return result, nil
}

func (service *KeyAnswerServiceImpl) UpdateAnswer(ctx echo.Context, request web.KeyUpdateAnswerReq, id int) (*domain.KeyAnswer, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAnswer, _ := service.KeyAnswerRepository.FindById(id)
	if existingAnswer == nil {
		return nil, fmt.Errorf("Answer Not Found")
	}

	answer := req.KeyUpdateAnswerRequestToKeyAnswerDomain(request)
	result, err := service.KeyAnswerRepository.Update(answer, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating Answer: %s", err.Error())
	}

	return result, nil
}

func (service *KeyAnswerServiceImpl) FindAnswerById(ctx echo.Context, id int) (*domain.KeyAnswer, error) {
	existingAnswer, _ := service.KeyAnswerRepository.FindById(id)
	if existingAnswer == nil {
		return nil, fmt.Errorf("Answer Not Found")
	}

	return existingAnswer, nil
}

func (service *KeyAnswerServiceImpl) FindAllAnswer(ctx echo.Context) ([]domain.KeyAnswer, error) {
	existingAnswer, _ := service.KeyAnswerRepository.FindAll()
	if existingAnswer == nil {
		return nil, fmt.Errorf("Answer Not Found")
	}

	return existingAnswer, nil
}

func (service *KeyAnswerServiceImpl) DeleteAnswer(ctx echo.Context, id int) error {
	existingAnswer, _ := service.KeyAnswerRepository.FindById(id)
	if existingAnswer == nil {
		return fmt.Errorf("Answer Not Found")
	}

	err := service.KeyAnswerRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting Answer: %s", err.Error())
	}

	return nil
}
