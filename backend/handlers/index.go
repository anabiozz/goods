package handlers

import (
	"net/http"

	"github.com/anabiozz/goods/backend/common"
)

// IndexHandler ...
func IndexHandler(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "./templates/index.html", nil)
	})
}
