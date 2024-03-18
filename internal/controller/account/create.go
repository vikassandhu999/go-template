package account

import (
	"betterx/internal/controller"
	"betterx/model"
	"betterx/pkg/logging"
	"errors"
	"net/http"
)

func (c *Controller) HandleCreate() http.Handler {
	type Form struct {
		Domain   string `json:"domain"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logging.FromContext(ctx).With("method", "module.HandleCreate")

		if r.Method == http.MethodGet {
			controller.Render(ctx, w, CreatePage(nil))
			return
		}

		form := Form{}
		if err := controller.BindForm(r, &form); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			controller.Render(ctx, w, CreatePage(err))
			return
		}

		_, err := c.db.CreateAccount(ctx, model.CreateAccountParams{
			Domain:   form.Domain,
			Email:    form.Email,
			Password: form.Password,
		})
		if err != nil {
			logger.Info("failed to create account", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			controller.Render(ctx, w, CreatePage(errors.New("internal server error")))
			return
		}
		http.Redirect(w, r, "/accounts/me", http.StatusSeeOther)
	})
}
