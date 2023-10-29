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

type QuizController interface {
	CreateQuizController(ctx echo.Context) error
	GetQuizByIdController(ctx echo.Context) error
	GetQuizByTitleController(ctx echo.Context) error
	GetAllQuizController(ctx echo.Context) error
	UpdateQuizController(ctx echo.Context) error
	DeleteQuizController(ctx echo.Context) error
}

type QuizControllerImpl struct {
	QuizService service.QuizService
}

func NewQuizController(quizService service.QuizService) QuizController {
	return &QuizControllerImpl{QuizService: quizService}
}

func (c *QuizControllerImpl) CreateQuizController(ctx echo.Context) error {
	QuizCreateRequest := web.QuizCreateRequest{}
	err := ctx.Bind(&QuizCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.QuizService.CreateQuiz(ctx, QuizCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "Quiz Already Exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Quiz Already Exist"))

		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Quiz Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Quiz", nil))

}

func (c *QuizControllerImpl) GetAllQuizController(ctx echo.Context) error {
	result, err := c.QuizService.FindAllQuiz(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Quiz Error"))
	}

	response := res.ConvertQuizResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Quiz", response))

}

func (c *QuizControllerImpl) GetQuizByIdController(ctx echo.Context) error {
	quizId := ctx.Param("id")
	quizIdInt, err := strconv.Atoi(quizId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.QuizService.FindQuizById(ctx, quizIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Quiz Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Quiz Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Quiz Error"))
	}

	response := res.FindQuizToQuizResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Quiz", response))
}

func (c *QuizControllerImpl) GetQuizByTitleController(ctx echo.Context) error {
	quizTitle := ctx.Param("title")

	result, err := c.QuizService.FindQuizByTitle(ctx, quizTitle)
	if err != nil {
		if strings.Contains(err.Error(), "Quiz Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Quiz Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Quiz Error"))
	}

	response := res.QuizDomainToQuizResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Quiz", response))
}

func (c *QuizControllerImpl) UpdateQuizController(ctx echo.Context) error {
	quizId := ctx.Param("id")
	quizIdInt, err := strconv.Atoi(quizId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	QuizUpdateRequest := web.QuizUpdateRequest{}
	err = ctx.Bind(&QuizUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.QuizService.UpdateQuiz(ctx, QuizUpdateRequest, quizIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "Quiz Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Quiz Not Found"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Quiz Error Internal"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Update Quiz", nil))
}

func (c *QuizControllerImpl) DeleteQuizController(ctx echo.Context) error {
	quizId := ctx.Param("id")
	quizIdInt, err := strconv.Atoi(quizId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.QuizService.DeleteQuiz(ctx, quizIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Quiz Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Quiz Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Quiz Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Delete Quiz", nil))
}
