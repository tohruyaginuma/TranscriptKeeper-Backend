package main

import (
	"context"
	"log/slog"

	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/registry"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/repository"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/route"
)

func main() {
	config.SetLogger()

	e := lib.SetEcho()

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

	e.Logger.Fatal(e.Start(":" + config.Port))
}
