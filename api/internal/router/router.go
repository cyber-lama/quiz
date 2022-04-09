package router

import (
	"api/internal/logger"
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
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

		var payload interface{}
		if r.Method == http.MethodPost {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Error(err)
			}
			payload = b
		} else if r.Method == http.MethodGet {
			payload = r.URL.Query().Encode()
		}

		log.Infof("--------------------------------")
		log.Infof("started %s %s, with payload: %s",
			r.Method, r.RequestURI, payload)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

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
		log.Infof("--------------------------------")
	})
}

func (r *Router) configureRouter() {
	r.Use(r.setRequestID)
	r.Use(r.logRequest)
}
