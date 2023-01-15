package ports

import "mmt.com/lolbank/domain"

type AccountRepo interface {
	AddAccount(acc domain.Account) domain.Account
	UpdateAccount(acc domain.Account) error
	GetAccount(accId string) (domain.Account, error)
	GetAccountByIdNumber(idNum string) (domain.Account, error)
}
