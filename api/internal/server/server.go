package server

import (
	"api/internal/configs"
	"api/internal/database"
	"api/internal/logger"
	"api/internal/router"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"strings"
)

type Server struct {
	configs *configs.Config
	router  *router.Router
	db      *database.Database
	logger  *logger.Logger

	servers []*HTTPServer
}

// Init a function that initializes
//all the important components of the application:
//routing, configuration from the env file,
//routing and processing http requests,
//connecting to the database
func Init() (*Server, error) {
	//// Get var from .env
	c := configs.Init()
	//// Plug logger
	l := logger.Init(c.LogLevel)
	//// Try to connect database
	//db, err := database.Connect(c.DBUrl)
	//if err != nil {
	//	return nil, err
	//}
	//// Init router with sum settings
	r := router.Init()
	//return &Server{
	//	configs: c,
	//	logger:  l,
	//	db:      db,
	//}, nil
	s := &Server{
		logger: l,
		router: r,
	}

	t, _ := net.Listen("tcp4", ":0")

	s.Accept(":3001", t)

	fmt.Printf("Addr: %s\n", t.Addr())
	err := s.serveAPI()
	if err != nil {
		return nil, err
	}
	return s, nil
}

type HTTPServer struct {
	srv *http.Server
	l   net.Listener
}

// Serve starts listening for inbound requests.
func (s *HTTPServer) Serve() error {
	return s.srv.Serve(s.l)
}

// Close closes the HTTPServer from listening for the inbound requests.
func (s *HTTPServer) Close() error {
	return s.l.Close()
}

func (s *Server) Accept(addr string, listeners ...net.Listener) {
	for _, listener := range listeners {
		httpServer := &HTTPServer{
			srv: &http.Server{
				Addr: addr,
			},
			l: listener,
		}
		s.servers = append(s.servers, httpServer)
	}
}
func (s *Server) Close() {
	for _, srv := range s.servers {
		if err := srv.Close(); err != nil {
			s.logger.Error(err)
		}
	}
}

//В чем разница между net.Listen и http.Server я понял
//net.Listen для такого: The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket"
//http.ListenAndServe для такого: HTTP-сервер.
//
//я в целом знаю что такое http и tcp
//tcp это более низкий уровень, но не понимаю почему мы не используем его, а используем http
//чтобы с сокетами работать нужно будет tcp или upd сервер поднимать ?
//Я так понял http использует tcp для связи между клиентом и сервером

func (s *Server) createMux() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode("test")
		if err != nil {
			return
		}
	})
}
func (s *Server) serveAPI() error {
	var chErrors = make(chan error, len(s.servers))
	for _, srv := range s.servers {
		srv.srv.Handler = s.createMux()
		go func(srv *HTTPServer) {
			var err error
			logrus.Infof("API listen on %s", srv.l.Addr())
			if err = srv.Serve(); err != nil && strings.Contains(err.Error(), "use of closed network connection") {
				err = nil
			}
			chErrors <- err
		}(srv)
	}
	for range s.servers {
		err := <-chErrors
		if err != nil {
			return err
		}
	}
	return nil
}
