package controllers

import (
	domain "TestTask-PGPro/internal/domain/repository"
	"TestTask-PGPro/internal/server/dto"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type LaunchController struct {
	launchRepository domain.ILaunchesRepository
	logger           slog.Logger
}

func NewLaunchController(logger slog.Logger, launchRepository domain.ILaunchesRepository) *LaunchController {
	return &LaunchController{
		launchRepository: launchRepository,

		logger: logger,
	}
}

func (c *LaunchController) LaunchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/launch" {
		if r.Method == http.MethodGet {
			c.getAllLaunches(w, r)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET or POST, got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
	} else {
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(w, "expect /command/<id>", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(pathParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodGet {
			c.getLaunch(w, r, id)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET or DELETE at /command/<id>, got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func (c *LaunchController) getAllLaunches(w http.ResponseWriter, r *http.Request) {
	launches, err := c.launchRepository.GetLaunches(r.Context())
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	response, err := json.Marshal(dto.LaunchesToLaunchesDTO(launches))
	if err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *LaunchController) getLaunch(w http.ResponseWriter, r *http.Request, id int) {
	launch, err := c.launchRepository.GetLaunch(r.Context(), id)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	response, err := json.Marshal(dto.Launch(launch))
	if err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
