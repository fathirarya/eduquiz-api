package controller

import (
	"eduquiz-api/model/web"
	"eduquiz-api/service"
	"eduquiz-api/utils/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type QuizResultController interface {
	CreateQuizResultController(ctx echo.Context) error
}

type QuizResultControllerImpl struct {
	QuizResultService    service.QuizResultService
	AttemptAnswerService service.AttemptAnswerService
}

func NewQuizResultController(quizResultService service.QuizResultService, attemptAnswerService service.AttemptAnswerService) QuizResultController {
	return &QuizResultControllerImpl{
		QuizResultService:    quizResultService,
		AttemptAnswerService: attemptAnswerService,
	}
}

func (c *QuizResultControllerImpl) CreateQuizResultController(ctx echo.Context) error {
	quizResultReq := web.QuizResultCreateRequest{}
	err := ctx.Bind(&quizResultReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.QuizResultService.PostResult(ctx, quizResultReq)
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
