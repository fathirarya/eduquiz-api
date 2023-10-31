package controller

import (
	"eduquiz-api/model/web"
	"eduquiz-api/service"
	"eduquiz-api/utils/helper"
	"eduquiz-api/utils/helper/middleware"
	"eduquiz-api/utils/res"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type TeacherController interface {
	RegisterTeacherController(ctx echo.Context) error
	LoginTeacherController(ctx echo.Context) error
	UpdateTeacherController(ctx echo.Context) error
	GetTeacherController(ctx echo.Context) error
	GetTeachersController(ctx echo.Context) error
	GetTeacherByNameController(ctx echo.Context) error
	DeleteTeacherController(ctx echo.Context) error
}

type TeacherControllerImpl struct {
	TeacherService service.TeacherService
}

func NewTeacherController(teacherService service.TeacherService) TeacherController {
	return &TeacherControllerImpl{TeacherService: teacherService}
}

func (c *TeacherControllerImpl) RegisterTeacherController(ctx echo.Context) error {
	teacherCreateRequest := web.TeacherCreateRequest{}
	err := ctx.Bind(&teacherCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.TeacherService.CreateTeacher(ctx, teacherCreateRequest)
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

func (c *TeacherControllerImpl) LoginTeacherController(ctx echo.Context) error {
	teacherLoginRequest := web.TeacherLoginRequest{}
	err := ctx.Bind(&teacherLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	response, err := c.TeacherService.LoginTeacher(ctx, teacherLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Invalid Email or Password") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Email or Password"))
		}
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign In Error"))
	}

	teacherLoginResponse := res.TeacherDomainToTeacherLoginResponse(response)

	token, err := middleware.GenerateTokenTeacher(response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}
	teacherLoginResponse.Token = token

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign In", teacherLoginResponse))
}

func (c *TeacherControllerImpl) UpdateTeacherController(ctx echo.Context) error {
	teacherId := ctx.Param("id")
	teacherIdInt, err := strconv.Atoi(teacherId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param id"))
	}

	teacherUpdateRequest := web.TeacherUpdateRequest{}
	err = ctx.Bind(&teacherUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.TeacherService.UpdateTeacher(ctx, teacherUpdateRequest, teacherIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Email Already Exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Update", nil))
}

func (c *TeacherControllerImpl) GetTeacherController(ctx echo.Context) error {
	teacherId := ctx.Param("id")
	teacherIdInt, err := strconv.Atoi(teacherId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param id"))
	}

	response, err := c.TeacherService.FindById(ctx, teacherIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Error"))
	}

	teacherResponse := res.TeacherDomainToTeacherResponse(response)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get", teacherResponse))
}

func (c *TeacherControllerImpl) GetTeachersController(ctx echo.Context) error {
	result, err := c.TeacherService.FindAll(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Error"))
	}

	teacherResponse := res.ConvertTeacherResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get", teacherResponse))
}

func (c *TeacherControllerImpl) GetTeacherByNameController(ctx echo.Context) error {
	teacherName := ctx.Param("username")

	result, err := c.TeacherService.FindByUsername(ctx, teacherName)
	if err != nil {
		if strings.Contains(err.Error(), "Teacher Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Teacher Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data By Name Error"))
	}

	teacherResponse := res.TeacherDomainToTeacherResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Teacher Data By Name", teacherResponse))
}

func (c *TeacherControllerImpl) DeleteTeacherController(ctx echo.Context) error {
	teacherId := ctx.Param("id")
	teacherIdInt, err := strconv.Atoi(teacherId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param id"))
	}

	err = c.TeacherService.DeleteTeacher(ctx, teacherIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Delete", nil))
}
