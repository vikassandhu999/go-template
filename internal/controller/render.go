package controller

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

type Template map[string]interface{}

func Render(ctx context.Context, w http.ResponseWriter, comp templ.Component) error {
	w.Header().Set("content-type", "text/html")
	return comp.Render(ctx, w)
}
