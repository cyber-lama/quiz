package basecontroller

import "net/http"

type IBase interface {
	Respond(w http.ResponseWriter, code int, data any)
	Message(w http.ResponseWriter, code int, data any)
	Error(w http.ResponseWriter, code int, err error)
}
