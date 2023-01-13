package services

import (
	"mmt.com/lolbank/domain"
	"mmt.com/lolbank/domain/requests"
	"mmt.com/lolbank/pkg/bankerrors"
	"mmt.com/lolbank/repos"
)

var (
	ErrorAlreadyExists     = bankerrors.NewBankError("an account already exists with the same id")
	ErrorInvalidAmount     = bankerrors.NewBankError("the given amount is not valid for a transaction")
	ErrorInsufficientFunds = bankerrors.NewBankError("no money")
)

type AccountService struct {
	repo *repos.AccountRepo
}

func NewAccountService() *AccountService {
	return &AccountService{
		repo: repos.NewAccountRepo(),
	}
}

func (svc *AccountService) CreateAccount(req requests.AccountCreateRequest) (domain.Account, error) {
	_, err := svc.repo.GetAccountByIdNumber(req.IdNumber)
	if err == nil {
		return domain.Account{}, ErrorAlreadyExists
	}

	newAccount := domain.Account{
		Balance:  req.InitialDeposit,
		IdNumber: req.IdNumber,
		Name:     req.Name,
	}

	return svc.repo.AddAccount(newAccount), nil
}

func (svc *AccountService) DepositMoney(accId string, req requests.AccountUpdateRequest) (domain.Account, error) {
	if req.Amount < 0 {
		return domain.Account{}, ErrorInvalidAmount
	}

	acc, err := svc.repo.GetAccount(accId)
	if err != nil {
		return domain.Account{}, err
	}

	acc.Balance = acc.Balance + req.Amount
	err = svc.repo.UpdateAccount(acc)
	if err != nil {
		return domain.Account{}, err
	}

	return acc, nil
}

func (svc *AccountService) WithdrawMoney(accId string, req requests.AccountUpdateRequest) (domain.Account, error) {
	if req.Amount < 0 {
		return domain.Account{}, ErrorInvalidAmount
	}

	acc, err := svc.repo.GetAccount(accId)
	if err != nil {
		return domain.Account{}, err
	}

	if req.Amount < acc.Balance {
		return domain.Account{}, ErrorInsufficientFunds
	}

	acc.Balance = acc.Balance - req.Amount
	err = svc.repo.UpdateAccount(acc)
	if err != nil {
		return domain.Account{}, err
	}

	return acc, nil
}

func (svc *AccountService) GetAccount(accId string) (domain.Account, error) {
	acc, err := svc.repo.GetAccount(accId)
	if err != nil {
		return domain.Account{}, err
	}

	return acc, nil
}
