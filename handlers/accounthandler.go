package handlers

import (
	"github.com/gin-gonic/gin"
	"mmt.com/lolbank/domain/requests"
	"mmt.com/lolbank/pkg/response"
	"mmt.com/lolbank/services"
	"net/http"
)

type AccountHandler struct {
	accountSvc *services.AccountService
}

func NewAccountHandler(accountSvc *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountSvc: accountSvc,
	}
}

func (hdl *AccountHandler) CreateAccount(ctx *gin.Context) {
	var request requests.AccountCreateRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	acc, err := hdl.accountSvc.CreateAccount(request)
	if err != nil {
		response.ProcessServiceError(ctx, err)
		return
	}

	response.SuccessResponse(ctx, http.StatusCreated, acc)
}

func (hdl *AccountHandler) DepositToAccount(ctx *gin.Context) {
	accountId := ctx.Param("accountId")
	var request requests.AccountUpdateRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	acc, err := hdl.accountSvc.DepositMoney(accountId, request)
	if err != nil {
		response.ProcessServiceError(ctx, err)
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, acc)
}

func (hdl *AccountHandler) WithdrawMoney(ctx *gin.Context) {
	accountId := ctx.Param("accountId")
	var request requests.AccountUpdateRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	acc, err := hdl.accountSvc.WithdrawMoney(accountId, request)
	if err != nil {
		response.ProcessServiceError(ctx, err)
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, acc)
}

func (hdl *AccountHandler) GetAccount(ctx *gin.Context) {
	accountId := ctx.Param("accountId")
	acc, err := hdl.accountSvc.GetAccount(accountId)
	if err != nil {
		response.ProcessServiceError(ctx, err)
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, acc)
}
