package main

import (
	"github.com/Futturi/testovoe/internal/handler"
	"github.com/Futturi/testovoe/internal/repository"
	"github.com/Futturi/testovoe/internal/server"
	"github.com/Futturi/testovoe/internal/service"
	"github.com/Futturi/testovoe/pkg"
	"log/slog"
	"os"
)

func main() {
	logg := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logg)
	cfg := pkg.Config{Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Hostname: os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBname:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE")}
	db, err := pkg.InitDB(cfg)
	if err != nil {
		slog.Error("error with initializing db", slog.Any("error", err))
	}
	repo := repository.NewRepository(db)
	servic := service.NewService(repo)
	han := handler.NewHandler(servic)
	if err := pkg.Migrat(os.Getenv("DB_HOST")); err != nil {
		slog.Error("error with migratedb", slog.Any("error", err))
	}
	serv := new(server.Server)
	slog.Info("server starting in", slog.String("port", os.Getenv("API_PORT")))
	if err = serv.InitServ(os.Getenv("API_PORT"), han.InitHandler()); err != nil {
		slog.Error("failed to run app", slog.Any("error", err))
	}
}
