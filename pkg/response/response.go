package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mmt.com/lolbank/pkg/bankerrors"
	"net/http"
)

type bankRes[T any] struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse[T any](ctx *gin.Context, status int, data T) {
	ctx.JSON(status, bankRes[T]{
		Success: true,
		Message: "Operation completed successfully.",
		Data:    data,
	})
}

func ErrorResponse(ctx *gin.Context, status int, message string) {
	ctx.AbortWithStatusJSON(status, bankRes[any]{
		Success: false,
		Message: message,
		Data:    nil,
	})
}

func ProcessServiceError(ctx *gin.Context, err error) {
	message := "Unknown error while processing request"

	var be *bankerrors.BankError
	if errors.As(err, &be) {
		ErrorResponse(ctx, http.StatusBadRequest, be.Error())
		return
	} else {
		ErrorResponse(ctx, http.StatusInternalServerError, message)
	}
}
