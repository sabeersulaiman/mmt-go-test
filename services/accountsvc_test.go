package services

import (
	"github.com/golang/mock/gomock"
	asserts "github.com/stretchr/testify/assert"
	"mmt.com/lolbank/domain"
	"mmt.com/lolbank/domain/requests"
	"mmt.com/lolbank/mocks"
	"mmt.com/lolbank/pkg/bankerrors"
	"testing"
)

func TestAccountService_CreateAccount_IfAccountExists_ThenReturnError(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()
	mockAccRepo := mocks.NewMockAccountRepo(cntrl)

	svc := NewAccountService(mockAccRepo)

	//setup
	mockAccRepo.EXPECT().
		GetAccountByIdNumber(gomock.Any()).
		Return(domain.Account{}, nil)

	// run
	_, err := svc.CreateAccount(requests.AccountCreateRequest{
		IdNumber: "some-id",
	})

	// assertions - tests
	assert := asserts.New(t)

	assert.NotNil(err)
	assert.ErrorIs(err, ErrorAlreadyExists)
}

func TestAccountService_CreateAccount_AndValid_TheAddAccount(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()
	mockAccRepo := mocks.NewMockAccountRepo(cntrl)

	svc := NewAccountService(mockAccRepo)

	//setup
	mockAccRepo.EXPECT().
		GetAccountByIdNumber(gomock.Any()).
		Return(domain.Account{}, bankerrors.ErrorNotFound)

	mockAccRepo.EXPECT().
		AddAccount(gomock.Any()).
		Return(domain.Account{
			AccountId: "uuid",
			IdNumber:  "some-id",
		})

	// run
	acc, err := svc.CreateAccount(requests.AccountCreateRequest{
		IdNumber: "some-id",
	})

	// assertions - tests
	assert := asserts.New(t)

	assert.Nil(err)
	assert.Equal("some-id", acc.IdNumber)
	assert.Equal("uuid", acc.AccountId)
}
