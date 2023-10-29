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

type TeacherService interface {
	CreateTeacher(ctx echo.Context, request web.TeacherCreateRequest) (*domain.Teacher, error)
	UpdateTeacher(ctx echo.Context, request web.TeacherUpdateRequest, id int) (*domain.Teacher, error)
	LoginTeacher(ctx echo.Context, request web.TeacherLoginRequest) (*domain.Teacher, error)
	FindById(ctx echo.Context, id int) (*domain.Teacher, error)
	FindAll(ctx echo.Context) ([]domain.Teacher, error)
	FindByUsername(ctx echo.Context, name string) (*domain.Teacher, error)
	DeleteTeacher(ctx echo.Context, id int) error
}

type TeacherServiceImpl struct {
	TeacherRepository repository.TeacherRepository
	Validate          *validator.Validate
}

func NewTeacherService(teacherRepository repository.TeacherRepository, validate *validator.Validate) *TeacherServiceImpl {
	return &TeacherServiceImpl{
		TeacherRepository: teacherRepository,
		Validate:          validate,
	}
}

func (service *TeacherServiceImpl) CreateTeacher(ctx echo.Context, request web.TeacherCreateRequest) (*domain.Teacher, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingTeacher, _ := service.TeacherRepository.FindByEmail(request.Email)
	if existingTeacher != nil {
		return nil, fmt.Errorf("Email Already Exist")
	}

	teacher := req.TeacherCreateRequestToTeacherDomain(request)

	teacher.Password = helper.HashPassword(teacher.Password)

	result, err := service.TeacherRepository.CreateTeacher(teacher)
	if err != nil {
		return nil, fmt.Errorf("Error when creating teacher: %s", err.Error())
	}

	return result, nil
}

func (service *TeacherServiceImpl) LoginTeacher(ctx echo.Context, request web.TeacherLoginRequest) (*domain.Teacher, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	exitingTeacher, err := service.TeacherRepository.FindByEmail(request.Email)
	fmt.Println(exitingTeacher)
	if err != nil {
		return nil, fmt.Errorf("Invalid Email")
	}

	teacher := req.TeacherLoginRequestToTeacherDomain(request)

	err = helper.ComparePassword(exitingTeacher.Password, teacher.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid Password")
	}

	return exitingTeacher, nil
}

func (service *TeacherServiceImpl) UpdateTeacher(ctx echo.Context, request web.TeacherUpdateRequest, id int) (*domain.Teacher, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	existingTeacher, _ := service.TeacherRepository.FindById(id)
	if existingTeacher != nil {
		return nil, fmt.Errorf("Teacher Not Found")
	}
	teacher := req.TeacherUpdateRequestToTeacherDomain(request)
	teacher.Password = helper.HashPassword(teacher.Password)

	result, err := service.TeacherRepository.UpdateTeacher(teacher, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating teacher: %s", err.Error())
	}

	return result, nil
}

func (service *TeacherServiceImpl) FindById(ctx echo.Context, id int) (*domain.Teacher, error) {

	exitingTeacher, err := service.TeacherRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("Teacher Not Found")
	}

	return exitingTeacher, nil
}

func (service *TeacherServiceImpl) FindAll(ctx echo.Context) ([]domain.Teacher, error) {

	teachers, err := service.TeacherRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Teachers Not Found")
	}

	return teachers, nil
}

func (service *TeacherServiceImpl) FindByUsername(ctx echo.Context, name string) (*domain.Teacher, error) {
	teacher, _ := service.TeacherRepository.FindByUsername(name)
	if teacher == nil {
		return nil, fmt.Errorf("Teacher Not Found")
	}

	return teacher, nil
}

func (service *TeacherServiceImpl) DeleteTeacher(ctx echo.Context, id int) error {
	existingTeacher, _ := service.TeacherRepository.FindById(id)
	if existingTeacher == nil {
		return fmt.Errorf("Teacher Not Found")
	}

	err := service.TeacherRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting teacher: %s", err.Error())
	}

	return nil
}
