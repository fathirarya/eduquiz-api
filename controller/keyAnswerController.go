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

type KeyAnswerController interface {
	CreateKeyAnswerController(ctx echo.Context) error
	GetKeyAnswerByIdController(ctx echo.Context) error
	GetAllKeyAnswerController(ctx echo.Context) error
	UpdateKeyAnswerController(ctx echo.Context) error
	DeleteKeyAnswerController(ctx echo.Context) error
}

type KeyAnswerControllerImpl struct {
	KeyAnswerService service.KeyAnswerService
}

func NewKeyAnswerController(keyAnswerService service.KeyAnswerService) KeyAnswerController {
	return &KeyAnswerControllerImpl{KeyAnswerService: keyAnswerService}
}

func (c *KeyAnswerControllerImpl) CreateKeyAnswerController(ctx echo.Context) error {
	KeyAnswerCreateRequest := web.KeyCreateAnswerReq{}
	err := ctx.Bind(&KeyAnswerCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.KeyAnswerService.CreateAnswer(ctx, KeyAnswerCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "KeyAnswer Already Exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("KeyAnswer Already Exist"))

		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("KeyAnswer Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create KeyAnswer", nil))
}

func (c *KeyAnswerControllerImpl) GetAllKeyAnswerController(ctx echo.Context) error {
	result, err := c.KeyAnswerService.FindAllAnswer(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("KeyAnswer Error"))
	}
	response := res.ConvertKeyAnswerResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All KeyAnswer", response))
}

func (c *KeyAnswerControllerImpl) GetKeyAnswerByIdController(ctx echo.Context) error {
	keyAnswerId := ctx.Param("id")
	keyAnswerIdInt, err := strconv.Atoi(keyAnswerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.KeyAnswerService.FindAnswerById(ctx, keyAnswerIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Key Answer Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Key Answer Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Key Answer Error"))
	}

	response := res.FindKeyAnswerToKeyAnswerResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Key Answer By Id", response))
}

func (c *KeyAnswerControllerImpl) UpdateKeyAnswerController(ctx echo.Context) error {
	keyAnswerId := ctx.Param("id")
	keyAnswerIdInt, err := strconv.Atoi(keyAnswerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	KeyAnswerUpdateRequest := web.KeyUpdateAnswerReq{}
	err = ctx.Bind(&KeyAnswerUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	_, err = c.KeyAnswerService.UpdateAnswer(ctx, KeyAnswerUpdateRequest, keyAnswerIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "Key Answer Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Key Answer Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Key Answer Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Update Key Answer", nil))
}

func (c *KeyAnswerControllerImpl) DeleteKeyAnswerController(ctx echo.Context) error {
	keyAnswerId := ctx.Param("id")
	keyAnswerIdInt, err := strconv.Atoi(keyAnswerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.KeyAnswerService.DeleteAnswer(ctx, keyAnswerIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Key Answer Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Key Answer Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Key Answer Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Delete Key Answer", nil))
}
