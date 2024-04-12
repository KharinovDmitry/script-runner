package controllers

import (
	"TestTask-PGPro/internal/domain/service"
	"TestTask-PGPro/internal/server/dto"
	"TestTask-PGPro/lib/byteconv"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type RunController struct {
	logger        slog.Logger
	launchService domain.ILaunchService
}

func NewRunController(logger slog.Logger, launchService domain.ILaunchService) *RunController {
	return &RunController{
		launchService: launchService,
		logger:        logger,
	}
}

func (c *RunController) RunHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	if len(pathParts) < 2 {
		http.Error(w, "expect /run/<id>", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		c.run(w, r, id)
	} else {
		http.Error(w, fmt.Sprintf("expect method POST at /run/<id>, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

func (c *RunController) run(w http.ResponseWriter, r *http.Request, id int) {
	launchId, err := c.launchService.Launch(r.Context(), id)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	w.Write(byteconv.Bytes(strconv.Itoa(launchId)))
}
