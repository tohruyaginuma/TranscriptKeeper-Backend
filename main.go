package main

import (
	"context"
	"log/slog"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/registry"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/repository"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/route"
)

const port = "8080"

func main() {
	config.SetLogger()

	e := config.SetEcho()

	cfg := config.Load()
	ctx := context.Background()
	db, err := repository.NewDB(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer db.Conn.Close()

	registry := registry.NewRegistry(db.Conn)
	route.SetRoute(e, registry)

	slog.Info("app starting")

	e.Logger.Fatal(e.Start(":" + port))
}
