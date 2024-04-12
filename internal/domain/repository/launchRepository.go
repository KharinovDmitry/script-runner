package domain

import (
	"TestTask-PGPro/internal/domain"
	"context"
)

type ILaunchesRepository interface {
	AddLaunch(ctx context.Context, commandID int) (int, error)
	AddOutputToLaunch(ctx context.Context, id int, output string) error
	GetLaunch(ctx context.Context, id int) (domain.Launch, error)
	GetLaunches(ctx context.Context) ([]domain.Launch, error)
}
