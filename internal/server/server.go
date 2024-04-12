package server

import (
	domain "TestTask-PGPro/internal/domain/service"
	"TestTask-PGPro/internal/server/controllers"
	"TestTask-PGPro/internal/storage"
	"log/slog"
	"net/http"
)

func MustRun(logger slog.Logger, launchService domain.ILaunchService, storage storage.Storage, address string) {
	commandController := controllers.NewCommandController(logger, storage.CommandsRepository)
	launchController := controllers.NewLaunchController(logger, storage.LaunchesRepository)

	RunController := controllers.NewRunController(logger, launchService)
	StopController := controllers.NewStopController(logger, launchService)

	mux := http.NewServeMux()

	mux.HandleFunc("/command", commandController.CommandHandler)
	mux.HandleFunc("/command/", commandController.CommandHandler)

	mux.HandleFunc("/launch", launchController.LaunchHandler)
	mux.HandleFunc("/launch/", launchController.LaunchHandler)

	mux.HandleFunc("/run/", RunController.RunHandler)
	mux.HandleFunc("/stop/", StopController.StopHandler)

	if err := http.ListenAndServe(address, mux); err != nil {
		panic(err.Error())
	}

}
