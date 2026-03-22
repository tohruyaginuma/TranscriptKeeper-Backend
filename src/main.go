package main

import (
	"context"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/infrastructure/db/postgres"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/middleware"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/registry"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/route"
)

func main() {
	config.SetLogger()
	cfg := config.Load()
	ctx := context.Background()
	db, err := postgres.NewDB(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer db.Conn.Close()

	registry := registry.NewRegistry(db.Conn)

	allowOrigins := []string{config.ClientURLWeb, config.ClientURLDesktop}
	e := config.SetEcho()
	e.Use(middleware.CORSMiddleware(allowOrigins))
	route.SetRoute(e, registry)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
