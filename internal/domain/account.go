package domain

import (
	"betterx/model"
	"betterx/pkg/aggregate"
)

type Account struct {
	aggregate.Root
	model.Account
}

type AccountCreated struct {
	Domain   string
	Email    string
	Password string
}

func (*AccountCreated) GetTopic() string {
	return "Account.Created"
}

type AccountUpdated struct {
	Domain   string
	Email    string
	Password string
}

func (*AccountUpdated) GetTopic() string {
	return "Account.Updated"
}

func NewAccount(domain string, email string, password string) (Account, error) {
	account := Account{Account: model.Account{Domain: domain, Email: email, Password: password}}
	account.AddEvent(&AccountCreated{
		Domain:   account.Domain,
		Email:    account.Email,
		Password: account.Password,
	})
	return account, nil
}
