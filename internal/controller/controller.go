package controller

import (
	"betterx/internal/controller/module"
	"betterx/internal/serverenv"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(ctx context.Context, env *serverenv.Env, r *mux.Router) error {
	sub := r.PathPrefix("").Subrouter()

	moduleController := module.New(env.DB())
	{
		sub := sub.PathPrefix("/modules").Subrouter()
		sub.Handle("", moduleController.HandleCreate()).Methods(http.MethodPost)
		sub.Handle("", moduleController.HandleIndex()).Methods(http.MethodGet)
		sub.Handle("/module/{id:[0-9]+}/edit", moduleController.HandleUpdate()).Methods(http.MethodGet)
		sub.Handle("/{id:[0-9]+}", moduleController.HandleShow()).Methods(http.MethodGet)
		sub.Handle("/{id:[0-9]+}", moduleController.HandleUpdate()).Methods(http.MethodPatch)
	}

	return nil
}
