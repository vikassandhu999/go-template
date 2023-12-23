package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

func BindForm(r *http.Request, data interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	decoder.SetAliasTag("form")

	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(data, r.PostForm); err != nil {
		return fmt.Errorf("failed to decode form: %w", err)
	}
	return nil
}
