package controller

import (
	"eduquiz-api/model/web"
	"eduquiz-api/service"
	"eduquiz-api/utils/helper"
	"eduquiz-api/utils/res"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type QuestionController interface {
	CreateQuestionController(ctx echo.Context) error
	UpdateQuestionController(ctx echo.Context) error
	GetQuestionByIdController(ctx echo.Context) error
	GetAllQuestionController(ctx echo.Context) error
	DeleteQuestionController(ctx echo.Context) error
}

type QuestionControllerImpl struct {
	QuestionService service.QuestionService
}

func NewQuestionController(questionService service.QuestionService) QuestionController {
	return &QuestionControllerImpl{QuestionService: questionService}
}

func (c *QuestionControllerImpl) CreateQuestionController(ctx echo.Context) error {
	questionCreateRequest := web.QuestionCreateRequest{}
	err := ctx.Bind(&questionCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.QuestionService.CreateQuestion(ctx, questionCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "Question Already Exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Question Already Exist"))

		}
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Question Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Question", nil))

}

func (c *QuestionControllerImpl) UpdateQuestionController(ctx echo.Context) error {
	questionId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Id"))
	}

	QuestionUpdateRequest := web.QuestionUpdateRequest{}
	err = ctx.Bind(&QuestionUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.QuestionService.UpdateQuestion(ctx, QuestionUpdateRequest, questionId)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Question Not Found") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Question Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Question Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Update Question", nil))
}

func (c *QuestionControllerImpl) GetQuestionByIdController(ctx echo.Context) error {
	questionId := ctx.Param("id")
	questionIdInt, err := strconv.Atoi(questionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.QuestionService.FindByIdQuestion(ctx, questionIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Question Not Found") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Question Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Question Error"))
	}

	response := res.QuestionDomainToQuestionResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Question Data By Id", response))
}

func (c *QuestionControllerImpl) GetAllQuestionController(ctx echo.Context) error {
	result, err := c.QuestionService.FindAllQuestion(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Question Error"))
	}

	response := res.ConvertQuestionResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Question Data", response))
}

func (c *QuestionControllerImpl) DeleteQuestionController(ctx echo.Context) error {
	questionId := ctx.Param("id")
	questionIdInt, err := strconv.Atoi(questionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.QuestionService.DeleteQuestion(ctx, questionIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Question Not Found") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Question Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Question Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Delete Question Data", nil))
}
