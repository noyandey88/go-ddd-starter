package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/noyandey/go-ddd-starter/config"
	"github.com/noyandey/go-ddd-starter/rest/handlers/user"
	"github.com/noyandey/go-ddd-starter/rest/middlewares"
)

type Server struct {
	cfg         *config.Config
	userHandler *user.Handler
}

func NewServer(cfg *config.Config, userHandler *user.Handler) *Server {
	return &Server{
		cfg:         cfg,
		userHandler: userHandler,
	}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.userHandler.RegisterRoutes(mux, manager)

	addr := fmt.Sprintf(":%d", server.cfg.HttpPort)

	fmt.Println("Server is running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}

}
