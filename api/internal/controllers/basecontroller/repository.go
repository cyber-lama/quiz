package basecontroller

import "net/http"

type Repository interface {
	Respond(w http.ResponseWriter, code int, data interface{})
	Message(w http.ResponseWriter, code int, data interface{})
	Error(w http.ResponseWriter, code int, err map[string]error)
}
