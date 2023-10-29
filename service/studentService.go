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

type StudentService interface {
	CreateStudent(ctx echo.Context, request web.StudentCreateRequest) (*domain.Student, error)
	LoginStudent(ctx echo.Context, request web.StudentLoginRequest) (*domain.Student, error)
	UpdateStudent(ctx echo.Context, request web.StudentUpdateRequest, id int) (*domain.Student, error)
	FindById(ctx echo.Context, id int) (*domain.Student, error)
	FindAll(ctx echo.Context) ([]domain.Student, error)
	FindByName(ctx echo.Context, name string) (*domain.Student, error)
	DeleteStudent(ctx echo.Context, id int) error
}

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
	Validate          *validator.Validate
}

func NewStudentService(studentRepository repository.StudentRepository, validate *validator.Validate) *StudentServiceImpl {
	return &StudentServiceImpl{
		StudentRepository: studentRepository,
		Validate:          validate,
	}
}

func (service *StudentServiceImpl) CreateStudent(ctx echo.Context, request web.StudentCreateRequest) (*domain.Student, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingStudent, _ := service.StudentRepository.FindByEmail(request.Email)
	if existingStudent != nil {
		return nil, fmt.Errorf("Email Already Exist")
	}

	student := req.StudentCreateRequestToStudentDomain(request)

	student.Password = helper.HashPassword(student.Password)

	result, err := service.StudentRepository.Create(student)
	if err != nil {
		return nil, fmt.Errorf("Error when creating student: %s", err.Error())
	}

	return result, nil
}

func (service *StudentServiceImpl) LoginStudent(ctx echo.Context, request web.StudentLoginRequest) (*domain.Student, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingStudent, err := service.StudentRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("Invalid Email or Password")
	}

	student := req.StudentLoginRequestToStudentDomain(request)

	err = helper.ComparePassword(existingStudent.Password, student.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid Email or Password")
	}

	return existingStudent, nil
}

func (service *StudentServiceImpl) UpdateStudent(ctx echo.Context, request web.StudentUpdateRequest, id int) (*domain.Student, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingStudent, _ := service.StudentRepository.FindById(id)
	if existingStudent == nil {
		return nil, fmt.Errorf("Student Not Found")
	}

	student := req.StudentUpdateRequestToStudentDomain(request)
	student.Password = helper.HashPassword(student.Password)

	result, err := service.StudentRepository.Update(student, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating student: %s", err.Error())
	}

	return result, nil
}

func (service *StudentServiceImpl) FindById(ctx echo.Context, id int) (*domain.Student, error) {

	existingStudent, _ := service.StudentRepository.FindById(id)
	if existingStudent == nil {
		return nil, fmt.Errorf("Student Not Found")
	}

	return existingStudent, nil
}

func (service *StudentServiceImpl) FindAll(ctx echo.Context) ([]domain.Student, error) {
	students, err := service.StudentRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Students Not Found")
	}

	return students, nil
}

func (service *StudentServiceImpl) FindByName(ctx echo.Context, name string) (*domain.Student, error) {
	student, _ := service.StudentRepository.FindByUsername(name)
	if student == nil {
		return nil, fmt.Errorf("Student Not Found")
	}

	return student, nil
}

func (context *StudentServiceImpl) DeleteStudent(ctx echo.Context, id int) error {

	existingStudent, _ := context.StudentRepository.FindById(id)
	if existingStudent == nil {
		return fmt.Errorf("Student Not Found")
	}

	err := context.StudentRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting student: %s", err)
	}

	return nil
}
