package module

import (
	"betterx/internal/controller"
	"betterx/model"
	"betterx/pkg/logging"
	"errors"
	"net/http"
)

func (c *Controller) HandleCreate() http.Handler {
	type Form struct {
		Domain   string `form:"domain"`
		Category string `form:"category"`
		Purpose  string `form:"purpose"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logging.FromContext(ctx).With("method", "module.HandleCreate")

		if r.Method == http.MethodGet {
			Create(nil).Render(ctx, w)
			return
		}

		form := Form{}
		if err := controller.BindForm(r, &form); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			Create(err).Render(ctx, w)
			return
		}

		db := model.New(c.db.Pool)

		_, err := db.CreateModule(ctx, model.CreateModuleParams{
			Domain:   form.Domain,
			Category: form.Category,
			Purpose:  form.Purpose,
		})
		if err != nil {
			logger.Error("failed to create module", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			Create(errors.New("internal server error")).Render(ctx, w)
			return
		}
		http.Redirect(w, r, "/modules", http.StatusSeeOther)
	})
}
