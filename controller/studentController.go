package controller

import (
	"eduquiz-api/model/web"
	"eduquiz-api/service"
	"eduquiz-api/utils/helper"
	"eduquiz-api/utils/helper/middleware"
	"eduquiz-api/utils/res"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type StudentController interface {
	RegisterStudentController(ctx echo.Context) error
	LoginStudentController(ctx echo.Context) error
	UpdateStudentController(ctx echo.Context) error
	GetStudentController(ctx echo.Context) error
	GetStudentsController(ctx echo.Context) error
	DeleteStudentController(ctx echo.Context) error
}

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentController(studentService service.StudentService) StudentController {
	return &StudentControllerImpl{StudentService: studentService}
}

func (c *StudentControllerImpl) RegisterStudentController(ctx echo.Context) error {
	studentCreateRequest := web.StudentCreateRequest{}
	err := ctx.Bind(&studentCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.StudentService.CreateStudent(ctx, studentCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "Email Already Exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", nil))
}

func (c *StudentControllerImpl) LoginStudentController(ctx echo.Context) error {
	studentLoginRequest := web.StudentLoginRequest{}
	err := ctx.Bind(&studentLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	response, err := c.StudentService.LoginStudent(ctx, studentLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Invalid Email or Password") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Email or Password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign In Error"))
	}

	studentLoginResponse := res.StudentDomainToStudentLoginResponse(response)

	token, err := middleware.GenerateTokenStudent(response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}
	studentLoginResponse.Token = token

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign In", studentLoginResponse))
}

func (c *StudentControllerImpl) GetStudentController(ctx echo.Context) error {
	studentId := ctx.Param("id")
	studentIdInt, err := strconv.Atoi(studentId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.StudentService.FindById(ctx, studentIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Students Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Students Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Student By ID Error"))
	}

	response := res.StudentDomainToStudentResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Student By Id", response))
}

func (c *StudentControllerImpl) GetStudentsController(ctx echo.Context) error {
	result, err := c.StudentService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Students Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Students Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Students Data Error"))
	}

	response := res.ConvertStudentResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Students Data", response))
}

func (c *StudentControllerImpl) UpdateStudentController(ctx echo.Context) error {
	studentId := ctx.Param("id")
	studentIdInt, err := strconv.Atoi(studentId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	studentUpdateRequest := web.StudentUpdateRequest{}
	err = ctx.Bind(&studentUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.StudentService.UpdateStudent(ctx, studentUpdateRequest, studentIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Student Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Student Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Student Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Updated Student Data", nil))
}

func (c *StudentControllerImpl) DeleteStudentController(ctx echo.Context) error {
	studentId := ctx.Param("id")
	studentIdInt, err := strconv.Atoi(studentId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.StudentService.DeleteStudent(ctx, studentIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Student Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Student Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Student Data Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Deleted Student Data", nil))
}
