package repository

import (
	"betterx/internal/domain"
	"betterx/model"
	"context"
	"fmt"
)

type Repo struct {
	db model.Queries
}

func (r *Repo) SaveAccount(ctx context.Context, account *domain.Account) error {
	events := account.PullEvent()
	for _, evt := range events {
		switch e := evt.(type) {
		case *domain.AccountCreated:
			_, err := r.db.CreateAccount(ctx, model.CreateAccountParams{
				Domain:   e.Domain,
				Email:    e.Email,
				Password: e.Password,
			})
			if err != nil {
				return fmt.Errorf("db.createaccount: %w", err)
			}
		case *domain.AccountUpdated:
			_, err := r.db.UpdateAccount(ctx, model.UpdateAccountParams{
				ID:       account.ID,
				Email:    e.Email,
				Password: e.Password,
			})
			if err != nil {
				return fmt.Errorf("db.createaccount: %w", err)
			}
		}
	}
	return nil
}
