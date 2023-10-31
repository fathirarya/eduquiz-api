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

type AttemptAnswerController interface {
	CreateAttemptAnswerController(ctx echo.Context) error
	GetAttemptAnswerByIdController(ctx echo.Context) error
	GetAllAttemptAnswerController(ctx echo.Context) error
}

type AttemptAnswerControllerImpl struct {
	AttempService service.AttemptAnswerService
}

func NewAttemptAnswerController(attempService service.AttemptAnswerService) AttemptAnswerController {
	return &AttemptAnswerControllerImpl{AttempService: attempService}
}

func (c *AttemptAnswerControllerImpl) CreateAttemptAnswerController(ctx echo.Context) error {
	attemptAnswerReq := web.AttemptAnswerCreateReq{}
	err := ctx.Bind(&attemptAnswerReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.AttempService.CreateAttemptAnswer(ctx, attemptAnswerReq)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "Question Already Exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Question Already Exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Question Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Question", nil))

}

func (c *AttemptAnswerControllerImpl) GetAttemptAnswerByIdController(ctx echo.Context) error {
	attemptAnswerId := ctx.Param("id")
	attemptAnswerIdInt, err := strconv.Atoi(attemptAnswerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Id"))
	}

	result, err := c.AttempService.FindByIdAttemptAnswer(ctx, attemptAnswerIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Answer with id") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Answer Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Answer Error"))

	}

	response := res.FindAttemptAnswerDomainToAttemptAnswerResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Answer", response))
}

func (c *AttemptAnswerControllerImpl) GetAllAttemptAnswerController(ctx echo.Context) error {
	result, err := c.AttempService.FindAllAttemptAnswer(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Answer with id") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Answer Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Answer Error"))

	}

	response := res.ConvertAttemptAnswerResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Answer", response))
}
