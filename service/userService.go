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

type UserService interface {
	CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.Users, error)
	LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.Users, error)
	UpdateUser(ctx echo.Context, request web.UserUpdateRequest, id int) (*domain.Users, error)
	FindById(ctx echo.Context, id int) (*domain.Users, error)
	FindAll(ctx echo.Context) ([]domain.Users, error)
	FindByName(ctx echo.Context, name string) (*domain.Users, error)
	DeleteUser(ctx echo.Context, id int) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.Users, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := service.UserRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("Email Already Exist")
	}

	user := req.UserCreateRequestToUserDomain(request)

	user.Password = helper.HashPassword(user.Password)

	result, err := service.UserRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("Error when creating user: %s", err.Error())
	}

	return result, nil
}

func (service *UserServiceImpl) LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.Users, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, err := service.UserRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("Invalid Email or Password")
	}

	user := req.UserLoginRequestToUserDomain(request)

	err = helper.ComparePassword(existingUser.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid Email or Password")
	}

	return existingUser, nil
}

func (service *UserServiceImpl) UpdateUser(ctx echo.Context, request web.UserUpdateRequest, id int) (*domain.Users, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := service.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("User Not Found")
	}

	user := req.UserUpdateRequestToUserDomain(request)
	user.Password = helper.HashPassword(user.Password)

	result, err := service.UserRepository.Update(user, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating user: %s", err.Error())
	}

	return result, nil
}

func (service *UserServiceImpl) FindById(ctx echo.Context, id int) (*domain.Users, error) {

	existingUser, _ := service.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("User Not Found")
	}

	return existingUser, nil
}

func (service *UserServiceImpl) FindAll(ctx echo.Context) ([]domain.Users, error) {
	users, err := service.UserRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Users Not Found")
	}

	return users, nil
}

func (service *UserServiceImpl) FindByName(ctx echo.Context, name string) (*domain.Users, error) {
	user, _ := service.UserRepository.FindByUsername(name)
	if user == nil {
		return nil, fmt.Errorf("User Not Found")
	}

	return user, nil
}

func (context *UserServiceImpl) DeleteUser(ctx echo.Context, id int) error {

	existingUser, _ := context.UserRepository.FindById(id)
	if existingUser == nil {
		return fmt.Errorf("User Not Found")
	}

	err := context.UserRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting user: %s", err)
	}

	return nil
}
