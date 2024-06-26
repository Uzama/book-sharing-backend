package server

import (
	"booking-management/utils/config"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	server  *http.Server
	address string
}

func NewHTTPServer(config config.Config, r *mux.Router) HTTPServer {

	address := config.App.Host + ":" + string(rune(config.App.Port))

	server := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,

		Handler: r,
	}

	httpServer := HTTPServer{
		server:  server,
		address: address,
	}

	return httpServer
}

func (srv HTTPServer) ListenAndServe(ctx context.Context) {

	log.Printf("server listening on %s", srv.address)

	err := srv.server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func (srv HTTPServer) Shutdown(ctx context.Context) {

	log.Println("stropping HTTP server")

	srv.server.SetKeepAlivesEnabled(false)

	err := srv.server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}
