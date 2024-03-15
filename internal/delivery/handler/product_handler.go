package handler

import (
	"net/http"

	"github.com/Dzikuri/shopifyx/internal/helper"
	"github.com/Dzikuri/shopifyx/internal/model"
	"github.com/Dzikuri/shopifyx/internal/usecase"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductUseCase *usecase.ProductUseCase
}

func NewProductHandler(productUseCase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		ProductUseCase: productUseCase,
	}
}

func (h *ProductHandler) ProductCreate(ctx echo.Context) error {
    tokenClaims := ctx.Get("user").(jwt.MapClaims)

    request := new(model.ProductCreateRequest)
    err := ctx.Bind(&request)
    if err != nil {
        return ctx.JSON(echo.ErrBadRequest.Code, map[string]interface{}{"message": "Bad Request", "error": err.Error()})
    }
    
    c := ctx.Request().Context()

    request.UserId = helper.GetUUID(tokenClaims["userId"].(string))

    response, err := h.ProductUseCase.ProductCreate(c, request)
    
    if err != nil {
        return ctx.JSON(echo.ErrBadRequest.Code, map[string]interface{}{"message": "Bad Request", "error": err.Error()})
    }

    return ctx.JSON(http.StatusOK, model.Response[model.ProductResponse] {
        Data: *response,
        Message: "Product created successfully",
    })
}