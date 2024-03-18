package routes

import (
	"betterx/internal/controller/account"
	"betterx/internal/serverenv"
	"betterx/model"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func Init(ctx context.Context, env *serverenv.Env, r *mux.Router) error {
	{
		sub := r.PathPrefix("").Subrouter()

		s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
		sub.PathPrefix("/static/").Handler(s)
	}

	db := model.New(env.DB().Pool)

	accountController := account.New(db)

	sub := r.PathPrefix("").Subrouter()
	{
		sub := sub.PathPrefix("/accounts").Subrouter()
		sub.Handle("/create", accountController.HandleCreate()).Methods(http.MethodGet)
		sub.Handle("/create", accountController.HandleCreate()).Methods(http.MethodPost)
	}

	return nil
}
