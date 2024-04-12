package postgres

import (
	"TestTask-PGPro/internal/domain"
	"TestTask-PGPro/internal/storage/dbModels"
	adapter "TestTask-PGPro/lib/adapter/db"
	"context"
)

type LaunchesRepository struct {
	db *adapter.PostgresAdapter
}

func NewLaunchesRepository(adapter *adapter.PostgresAdapter) *LaunchesRepository {
	return &LaunchesRepository{db: adapter}
}

func (l *LaunchesRepository) AddLaunch(ctx context.Context, commandID int) (int, error) {
	sql := `insert into launches(command_id) values ($1) returning id`
	var id int
	err := l.db.ExecuteAndGet(ctx, &id, sql, commandID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (l *LaunchesRepository) AddOutputToLaunch(ctx context.Context, id int, output string) error {
	sql := `update launches set output = concat(output, text($1)) where id = $2`

	return l.db.Execute(ctx, sql, output, id)
}

func (l *LaunchesRepository) GetLaunch(ctx context.Context, id int) (domain.Launch, error) {
	sql := `select id, command_id, output from launches where id = $1`

	var launch dbModels.Launch
	err := l.db.ExecuteAndGet(ctx, &launch, sql, id)

	if err != nil {
		return domain.Launch{}, err
	}

	return domain.Launch(launch), nil
}

func (l *LaunchesRepository) GetLaunches(ctx context.Context) ([]domain.Launch, error) {
	sql := `select id, command_id, output from launches`

	var launch []dbModels.Launch
	err := l.db.Query(ctx, &launch, sql)

	if err != nil {
		return nil, err
	}

	return dbModels.DBLaunchesToLaunches(launch), nil
}
