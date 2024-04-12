package controllers

import (
	domain "TestTask-PGPro/internal/domain/repository"
	"TestTask-PGPro/internal/server/dto"
	"TestTask-PGPro/lib/byteconv"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type CommandController struct {
	commandRepository domain.ICommandsRepository
	logger            slog.Logger
}

func NewCommandController(logger slog.Logger, commandRepository domain.ICommandsRepository) *CommandController {
	return &CommandController{
		commandRepository: commandRepository,

		logger: logger,
	}
}

func (c *CommandController) CommandHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/command" {
		if r.Method == http.MethodPost {
			c.createCommandHandler(w, r)
		} else if r.Method == http.MethodGet {
			c.getAllCommandsHandler(w, r)
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

		if r.Method == http.MethodDelete {
			c.deleteCommandHandler(w, r, id)
		} else if r.Method == http.MethodGet {
			c.getCommandHandler(w, r, id)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET or DELETE at /command/<id>, got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func (c *CommandController) createCommandHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.AddCommandRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := c.commandRepository.AddCommand(r.Context(), req.Text)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	w.Write(byteconv.Bytes(strconv.Itoa(id)))
}

func (c *CommandController) getAllCommandsHandler(w http.ResponseWriter, r *http.Request) {
	commands, err := c.commandRepository.GetCommands(r.Context())
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	response, err := json.Marshal(dto.CommandsToCommandsDTO(commands))
	if err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *CommandController) getCommandHandler(w http.ResponseWriter, r *http.Request, id int) {
	command, err := c.commandRepository.GetCommand(r.Context(), id)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	response, err := json.Marshal(dto.Command(command))
	if err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *CommandController) deleteCommandHandler(w http.ResponseWriter, r *http.Request, id int) {
	err := c.commandRepository.DeleteCommand(r.Context(), id)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	w.WriteHeader(http.StatusOK)
}
