package dto

import "TestTask-PGPro/internal/domain"

type Command struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func CommandsToCommandsDTO(commands []domain.Command) []Command {
	res := make([]Command, len(commands))
	for i, _ := range res {
		res[i] = Command((commands[i]))
	}
	return res
}
