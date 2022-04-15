package basecontroller

import (
	"net/http"
)

type Repository interface {
	Respond[D any](w http.ResponseWriter, code int, data []D)
	Message[D any](w http.ResponseWriter, code int, data []D)
	Error(w http.ResponseWriter, code int, err error)
}
