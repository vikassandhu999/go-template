package serverenv

import (
	"betterx/pkg/mongodb"
	"betterx/pkg/postgres"
	"context"
)

type Env struct {
	db      *postgres.DB
	mongoDB *mongodb.DB
}

type Option func(*Env) *Env

func New(opts ...Option) *Env {
	env := &Env{}
	for _, opt := range opts {
		env = opt(env)
	}
	return env
}

func WithDB(db *postgres.DB) Option {
	return func(e *Env) *Env {
		e.db = db
		return e
	}
}

func (e *Env) DB() *postgres.DB {
	return e.db
}

func WithMongo(db *mongodb.DB) Option {
	return func(e *Env) *Env {
		e.mongoDB = db
		return e
	}
}

func (e *Env) Mongo() *mongodb.DB {
	return e.mongoDB
}

func (e *Env) Close(ctx context.Context) error {
	e.db.Close(ctx)
	return nil
}
