package repos

import (
	"github.com/google/uuid"
	"mmt.com/lolbank/domain"
	"mmt.com/lolbank/pkg/bankerrors"
)

type AccountRepo struct {
	store map[string]domain.Account
}

func NewAccountRepo() *AccountRepo {
	return &AccountRepo{
		store: make(map[string]domain.Account),
	}
}

func (r *AccountRepo) AddAccount(acc domain.Account) domain.Account {
	newAccountId := uuid.New().String()
	acc.AccountId = newAccountId
	r.store[acc.AccountId] = acc

	return acc
}

func (r *AccountRepo) UpdateAccount(acc domain.Account) error {
	existing := r.getAccount(acc.AccountId)
	if existing == nil {
		return bankerrors.ErrorNotFound
	}

	existing.Balance = acc.Balance
	r.store[acc.AccountId] = *existing
	return nil
}

func (r *AccountRepo) GetAccount(accId string) (domain.Account, error) {
	existing := r.getAccount(accId)
	if existing == nil {
		return domain.Account{}, bankerrors.ErrorNotFound
	}

	return *existing, nil
}

func (r *AccountRepo) GetAccountByIdNumber(idNum string) (domain.Account, error) {
	for _, v := range r.store {
		if v.IdNumber == idNum {
			return v, nil
		}
	}

	return domain.Account{}, bankerrors.ErrorNotFound
}

func (r *AccountRepo) getAccount(accId string) *domain.Account {
	if acc, ok := r.store[accId]; ok {
		return &acc
	}

	return nil
}
