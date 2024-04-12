package dbModels

import "TestTask-PGPro/internal/domain"

type Command struct {
	ID   int    `db:"id"`
	Text string `db:"text"`
}

func DBCommandsToCommands(commands []Command) []domain.Command {
	res := make([]domain.Command, len(commands))
	for i, _ := range res {
		res[i] = domain.Command(commands[i])
	}
	return res
}
