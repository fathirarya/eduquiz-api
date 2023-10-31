package controller

import (
	"eduquiz-api/model/web"
	"eduquiz-api/service"
	"eduquiz-api/utils/helper"
	"eduquiz-api/utils/res"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type QuizCategoryController interface {
	CreateQuizCategoryController(ctx echo.Context) error
	GetQuizCategoryByIdController(ctx echo.Context) error
	GetAllQuizCategoryController(ctx echo.Context) error
	DeleteQuizCategoryController(ctx echo.Context) error
}

type QuizCategoryControllerImpl struct {
	QuizCategoryService service.QuizCategoryService
}

func NewQuizCategoryController(quizCategoryService service.QuizCategoryService) QuizCategoryController {
	return &QuizCategoryControllerImpl{QuizCategoryService: quizCategoryService}
}

func (c *QuizCategoryControllerImpl) CreateQuizCategoryController(ctx echo.Context) error {
	quizCategoryCreateReq := web.QuizCategoryCreateRequest{}
	err := ctx.Bind(&quizCategoryCreateReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.QuizCategoryService.CreateQuizCategory(ctx, quizCategoryCreateReq)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "Category Already Exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Category Already Exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Category Error"))
	}


	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Category", nil))


}

func (c *QuizCategoryControllerImpl) GetAllQuizCategoryController(ctx echo.Context) error {
	result, err := c.QuizCategoryService.FindAllQuizCategory(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Category Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Category Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Category Data Error"))
	}

	response := res.ConvertQuizCategoryResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Category", response))
}

func (c *QuizCategoryControllerImpl) GetQuizCategoryByIdController(ctx echo.Context) error {
	quizCategoryId := ctx.Param("id")
	quizCategoryIdInt, err := strconv.Atoi(quizCategoryId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.QuizCategoryService.FindQuizCategoryById(ctx, quizCategoryIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Category Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Category Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Category Data Error"))
	}

	response := res.QuizCategoryDomainToQuizCategoryResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Category", response))
}

func (c *QuizCategoryControllerImpl) DeleteQuizCategoryController(ctx echo.Context) error {
	quizCategoryId := ctx.Param("id")
	quizCategoryIdInt, err := strconv.Atoi(quizCategoryId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.QuizCategoryService.DeleteQuizCategory(ctx, quizCategoryIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Category Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Category Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Category Data Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Delete Category Data", nil))
}
