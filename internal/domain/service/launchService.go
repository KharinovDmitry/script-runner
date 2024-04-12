package domain

import "context"

type ILaunchService interface {
	Launch(ctx context.Context, commandId int) (int, error)
	Stop(ctx context.Context, launchId int) error
}
