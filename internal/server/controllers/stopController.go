package controllers

import (
	domain "TestTask-PGPro/internal/domain/service"
	"TestTask-PGPro/internal/server/dto"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type StopController struct {
	logger        slog.Logger
	launchService domain.ILaunchService
}

func NewStopController(logger slog.Logger, launchService domain.ILaunchService) *StopController {
	return &StopController{
		launchService: launchService,
		logger:        logger,
	}
}

func (c *StopController) StopHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	if len(pathParts) < 2 {
		http.Error(w, "expect /stop/<id>", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		c.stop(w, r, id)
	} else {
		http.Error(w, fmt.Sprintf("expect method POST at /stop/<id>, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

func (c *StopController) stop(w http.ResponseWriter, r *http.Request, id int) {
	err := c.launchService.Stop(r.Context(), id)
	if err != nil {
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}
	w.WriteHeader(http.StatusOK)
}
