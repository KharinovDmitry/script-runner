package app

import (
	"TestTask-PGPro/cmd/migrator"
	"TestTask-PGPro/internal/config"
	"TestTask-PGPro/internal/server"
	"TestTask-PGPro/internal/service"
	"TestTask-PGPro/internal/storage"
	adapter "TestTask-PGPro/lib/adapter/db"
	"TestTask-PGPro/lib/adapter/executor"
	"context"
	"errors"
	"log/slog"
	"os"
)

func MustRun(cfg *config.Config) {
	logger, err := setupLogger(cfg.Env)
	if err != nil {
		panic(err.Error())
	}

	migrator.MustRun(cfg.DriverName, cfg.ConnStr, cfg.MigrationsDir)
	db := adapter.NewPostgresAdapter(cfg.TimeoutDB)
	db.Connect(context.Background(), cfg.ConnStr)
	defer db.Close()

	store := storage.NewStorage(db)
	store.InitStorage()

	executor := executor.NewLinuxAdapter()

	launchService := service.NewLaunchService(*logger, executor, store.CommandsRepository, store.LaunchesRepository)

	server.MustRun(*logger, launchService, *store, cfg.Address)
}

func setupLogger(env string) (*slog.Logger, error) {
	switch env {
	case "local":
		return slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			})), nil
	case "dev":
		file, err := os.OpenFile("app_logs", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}
		return slog.New(slog.NewTextHandler(
			file,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			})), nil
	case "prod":
		file, err := os.OpenFile("app_logs", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}
		return slog.New(slog.NewTextHandler(
			file,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			})), nil
	default:
		return nil, errors.New("Unknown ENV: " + env)
	}
}
