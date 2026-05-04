// package rest

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/noyandey/go-ddd-starter/config"
// 	"github.com/noyandey/go-ddd-starter/rest/handlers/user"
// 	"github.com/noyandey/go-ddd-starter/rest/middlewares"
// )

// type Server struct {
// 	cfg         *config.Config
// 	userHandler *user.Handler
// }

// func NewServer(cfg *config.Config, userHandler *user.Handler) *Server {
// 	return &Server{
// 		cfg:         cfg,
// 		userHandler: userHandler,
// 	}
// }

// func (server *Server) Start() {
// 	manager := middlewares.NewManager()
// 	manager.Use(
// 		middlewares.Preflight,
// 		middlewares.Cors,
// 		middlewares.Logger,
// 	)

// 	mux := http.NewServeMux()
// 	wrappedMux := manager.WrapMux(mux)

// 	server.userHandler.RegisterRoutes(mux, manager)

// 	addr := fmt.Sprintf(":%d", server.cfg.HttpPort)

// 	fmt.Println("Server is running on port", addr)
// 	err := http.ListenAndServe(addr, wrappedMux)

// 	if err != nil {
// 		fmt.Println("Error starting server", err)
// 		os.Exit(1)
// 	}

// }

// package rest

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/danielgtaylor/huma/v2"
// 	"github.com/danielgtaylor/huma/v2/adapters/humago"
// 	"github.com/noyandey/go-ddd-starter/config"
// 	"github.com/noyandey/go-ddd-starter/rest/handlers/user"
// 	"github.com/noyandey/go-ddd-starter/rest/middlewares"
// )

// type Server struct {
// 	cfg         *config.Config
// 	userHandler *user.Handler
// 	api         huma.API // 👈 add this
// }

// func NewServer(cfg *config.Config, userHandler *user.Handler) *Server {
// 	return &Server{
// 		cfg:         cfg,
// 		userHandler: userHandler,
// 	}
// }

// func (server *Server) Start() {
// 	manager := middlewares.NewManager()
// 	manager.Use(
// 		middlewares.Preflight,
// 		middlewares.Cors,
// 		middlewares.Logger,
// 	)

// 	mux := http.NewServeMux()
// 	wrappedMux := manager.WrapMux(mux)

// 	// 👇 Add Huma on top of your existing mux
// 	api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))
// 	server.api = api

// 	// Pass api to your handlers for Huma route registration
// 	server.userHandler.RegisterRoutes(mux, manager)

// 	addr := fmt.Sprintf(":%d", server.cfg.HttpPort)
// 	fmt.Println("Server is running on port", addr)
// 	err := http.ListenAndServe(addr, wrappedMux)
// 	if err != nil {
// 		fmt.Println("Error starting server", err)
// 		os.Exit(1)
// 	}
// }

// package rest

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/danielgtaylor/huma/v2"
// 	"github.com/danielgtaylor/huma/v2/adapters/humago" // ✅ correct
// 	"github.com/noyandey/go-ddd-starter/config"
// 	"github.com/noyandey/go-ddd-starter/rest/handlers/user"
// 	"github.com/noyandey/go-ddd-starter/rest/middlewares"
// )

// type Server struct {
// 	cfg         *config.Config
// 	userHandler *user.Handler
// 	api         huma.API
// }

// func NewServer(cfg *config.Config, userHandler *user.Handler) *Server {
// 	return &Server{
// 		cfg:         cfg,
// 		userHandler: userHandler,
// 	}
// }

// func (server *Server) Start() {
// 	manager := middlewares.NewManager()
// 	manager.Use(
// 		middlewares.Preflight,
// 		middlewares.Cors,
// 		middlewares.Logger,
// 	)

// 	mux := http.NewServeMux()
// 	wrappedMux := manager.WrapMux(mux)

// 	// ✅ correct adapter
// 	config := huma.DefaultConfig("My API", "1.0.0")
// 	config.DocsRenderer = huma.DocsRendererSwaggerUI // ✅ switch to Swagger UI

// 	api := humago.New(mux, config)
// 	server.api = api
// 	// ✅ pass api down to RegisterRoutes
// 	server.userHandler.RegisterRoutes(mux, manager)

// 	addr := fmt.Sprintf(":%d", server.cfg.HttpPort)
// 	fmt.Println("Server is running on port", addr)
// 	err := http.ListenAndServe(addr, wrappedMux)
// 	if err != nil {
// 		fmt.Println("Error starting server", err)
// 		os.Exit(1)
// 	}
// }

package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/noyandey/go-ddd-starter/config"
	"github.com/noyandey/go-ddd-starter/rest/handlers/user"
	"github.com/noyandey/go-ddd-starter/rest/middlewares"
)

// --- Demo structs ---

type GreetInput struct {
	Name string `path:"name" doc:"Name to greet"`
}

type GreetOutput struct {
	Body struct {
		Message string `json:"message" doc:"Greeting message"`
	}
}

type PingOutput struct {
	Body struct {
		Status string `json:"status" doc:"Server status"`
	}
}

// --------------------

type Server struct {
	cfg         *config.Config
	userHandler *user.Handler
	api         huma.API
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

	// ✅ Swagger UI config
	// cfg := huma.DefaultConfig("My API", "1.0.0")
	// cfg.DocsRenderer = huma.DocsRendererSwaggerUI
	// cfg.Info.Description = "My API documentation"

	// api := humago.New(mux, cfg)
	// server.api = api

	cfg := huma.DefaultConfig("My API", "1.0.0")

	// ✅ Swagger UI
	cfg.DocsRenderer = huma.DocsRendererSwaggerUI

	// ✅ API Info
	cfg.Info.Description = "My API documentation"
	cfg.Info.TermsOfService = "https://myapi.com/terms"
	cfg.Info.Contact = &huma.Contact{
		Name:  "My Team",
		Email: "support@myapi.com",
		URL:   "https://myapi.com",
	}
	cfg.Info.License = &huma.License{
		Name: "MIT",
		URL:  "https://opensource.org/licenses/MIT",
	}

	// ✅ Servers
	cfg.Servers = []*huma.Server{
		{URL: "http://localhost:4000", Description: "Local development"},
		{URL: "https://api.myapp.com", Description: "Production"},
	}

	// ✅ Bearer token security
	cfg.Components = &huma.Components{
		SecuritySchemes: map[string]*huma.SecurityScheme{
			"bearerAuth": {
				Type:         "http",
				Scheme:       "bearer",
				BearerFormat: "JWT",
				Description:  "Enter your JWT bearer token",
			},
		},
	}

	// ✅ Apply security globally
	cfg.Security = []map[string][]string{
		{"bearerAuth": {}},
	}

	// ✅ plain New — no prefix
	api := humago.New(mux, cfg)
	server.api = api

	// ✅ Demo route 1 — GET /ping
	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/ping",
		Summary:     "Health check",
		Description: "Returns pong if server is running",
		Tags:        []string{"Health"},
	}, func(ctx context.Context, input *struct{}) (*PingOutput, error) {
		resp := &PingOutput{}
		resp.Body.Status = "pong"
		return resp, nil
	})

	// ✅ Demo route 2 — GET /greet/{name}
	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/greet/{name}",
		Summary:     "Greet a user",
		Description: "Returns a greeting message for the given name",
		Tags:        []string{"Health"},
	}, func(ctx context.Context, input *GreetInput) (*GreetOutput, error) {
		resp := &GreetOutput{}
		resp.Body.Message = "Hello, " + input.Name + "!"
		return resp, nil
	})

	server.userHandler.RegisterRoutes(mux, manager)

	addr := fmt.Sprintf(":%d", server.cfg.HttpPort)
	fmt.Println("Server is running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}
}
