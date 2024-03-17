package server

import (
	"net/http"
	"time"

	"github.com/trevatk/web3.5/internal/adapter/setup"
)

// New return new http server
func New(cfg *setup.Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":" + cfg.Server.HTTPPort,
		Handler:      handler,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}
