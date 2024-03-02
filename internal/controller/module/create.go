package module

import (
	"net/http"
)

func (c *Controller) HandleCreate() http.Handler {
	type Form struct {
		Domain   string `form:"domain"`
		Category string `form:"category"`
		Purpose  string `form:"purpose"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		// logger := logging.FromContext(ctx).With("method", "module.HandleCreate")

		// if r.Method == http.MethodGet {
		// 	controller.Render(ctx, w, CreatePage(nil))
		// 	return
		// }

		// form := Form{}
		// if err := controller.BindForm(r, &form); err != nil {
		// 	w.WriteHeader(http.StatusUnprocessableEntity)
		// 	w.Write([]byte(err.Error()))
		// 	controller.Render(ctx, w, CreatePage(err))
		// 	return
		// }

		// db := model.New(c.db.Pool)

		// _, err := db.CreateModule(ctx, model.CreateModuleParams{
		// 	Domain:   form.Domain,
		// 	Category: form.Category,
		// 	Purpose:  form.Purpose,
		// })
		// if err != nil {
		// 	logger.Error("failed to create module", "error", err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	controller.Render(ctx, w, CreatePage(errors.New("internal server error")))
		// 	return
		// }
		http.Redirect(w, r, "/modules", http.StatusSeeOther)
	})
}
