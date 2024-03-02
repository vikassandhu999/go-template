package module

import "betterx/model"

type Controller struct {
	db *model.Queries
}

func New(db *model.Queries) *Controller {
	return &Controller{db: db}
}
