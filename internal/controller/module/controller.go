package module

import (
	"betterx/pkg/postgres"
)

type Controller struct {
	db *postgres.DB
}

func New(db *postgres.DB) *Controller {
	return &Controller{db: db}
}
