package router

import (
	"api/internal/logger"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Router struct {
	*mux.Router
}

func Init() *Router {
	r := &Router{mux.NewRouter()}
	r.configureRouter()
	return r
}

const (
	ctxKeyRequestID = iota
)

func (r *Router) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (r *Router) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.Init("trace")
		log.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		log.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var reqBody interface{}
		err := json.NewDecoder(r.Body).Decode(reqBody)
		if err != nil {
			log.Error(err)
		} else {
			log.Info(reqBody)
		}
		log.Info(r.Method)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		log.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (r *Router) configureRouter() {
	r.Use(r.setRequestID)
	r.Use(r.logRequest)
	r.HandleFunc("/test", r.handleUsersCreate()).Methods("POST")
}

func (r *Router) respond(w http.ResponseWriter, req *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}

func (r *Router) handleUsersCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		r.respond(w, req, http.StatusCreated, "test")
	}
}
