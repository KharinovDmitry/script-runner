package domain

import (
	"TestTask-PGPro/internal/domain"
	"context"
)

type ICommandsRepository interface {
	AddCommand(ctx context.Context, text string) (int, error)
	DeleteCommand(ctx context.Context, id int) error
	GetCommand(ctx context.Context, id int) (domain.Command, error)
	GetCommands(ctx context.Context) ([]domain.Command, error)
}
