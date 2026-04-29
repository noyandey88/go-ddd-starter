package cmd

import (
	"fmt"
	"os"

	"github.com/noyandey/go-ddd-starter/config"
	"github.com/noyandey/go-ddd-starter/infra/db"
	"github.com/noyandey/go-ddd-starter/internal/user"
	"github.com/noyandey/go-ddd-starter/repo"
	"github.com/noyandey/go-ddd-starter/rest"
	userHandler "github.com/noyandey/go-ddd-starter/rest/handlers/user"
)

func Serve() {
	cfg := config.GetConfig()

	dbCon, err := db.NewConnection(cfg.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	userService := user.NewService(userRepo)

	//middlewares
	// middlewares := middlewares.NewMiddlewares(cfg)

	// handlers
	userHandler := userHandler.NewHandler(userService, cfg)

	server := rest.NewServer(cfg, userHandler)

	server.Start()
}
