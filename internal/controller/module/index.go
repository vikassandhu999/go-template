package module

import (
	"betterx/internal/controller"
	"betterx/model"
	"betterx/pkg/logging"
	"errors"
	"net/http"
)

func (c *Controller) HandleIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logging.FromContext(ctx).With("method", "module.HandleCreate")

		db := model.New(c.db.Pool)

		modules, err := db.ListModules(ctx, model.ListModulesParams{Limit: 10, Offset: 0})
		if err != nil {
			logger.Error("failed to list modules", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			controller.Render(ctx, w, Index([]model.Module{}, errors.New("internal server error")))
			return
		}
		controller.Render(ctx, w, Index(modules, nil))
	})
}
