package module

import "net/http"

func (c *Controller) HandleUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
