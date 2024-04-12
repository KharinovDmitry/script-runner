package service

import (
	domain "TestTask-PGPro/internal/domain/repository"
	"TestTask-PGPro/lib/adapter/executor"
	"TestTask-PGPro/lib/byteconv"
	"context"
	"github.com/pkg/errors"
	"log/slog"
	"sync"
)

type LaunchService struct {
	logger             slog.Logger
	executor           executor.LinuxAdapter
	commandsRepository domain.ICommandsRepository
	launchesRepository domain.ILaunchesRepository

	runs sync.Map
}

func NewLaunchService(logger slog.Logger, executor executor.LinuxAdapter, commandsRepository domain.ICommandsRepository, launchesRepository domain.ILaunchesRepository) *LaunchService {
	return &LaunchService{
		logger:             logger,
		executor:           executor,
		commandsRepository: commandsRepository,
		launchesRepository: launchesRepository,
	}
}

func (l *LaunchService) Launch(ctx context.Context, commandId int) (int, error) {
	command, err := l.commandsRepository.GetCommand(ctx, commandId)
	if err != nil {
		return 0, err
	}

	launchId, err := l.launchesRepository.AddLaunch(ctx, commandId)
	if err != nil {
		return 0, err
	}

	go func() {
		output := make(chan []byte)

		go func() {
			runCtx, stop := context.WithCancel(context.Background())
			l.runs.Store(launchId, stop)

			if err := l.executor.Run(runCtx, command.Text, output); err != nil {
				l.logger.Error(err.Error())
			}
		}()

		for {
			res, ok := <-output
			if ok {
				err = l.launchesRepository.AddOutputToLaunch(ctx, launchId, byteconv.String(res))
				if err != nil {
					l.logger.Error(err.Error())
				}
			} else {
				break
			}
		}
	}()

	return launchId, nil
}

func (l *LaunchService) Stop(ctx context.Context, commandId int) error {
	v, ok := l.runs.LoadAndDelete(commandId)
	if !ok {
		return ErrNotFound
	}
	stop, ok := v.(context.CancelFunc)
	if !ok {
		return errors.Wrapf(ErrStop, ": commandID - %d", commandId)
	}
	stop()
	return nil
}
