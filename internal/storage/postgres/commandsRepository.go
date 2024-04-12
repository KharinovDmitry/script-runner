package postgres

import (
	"TestTask-PGPro/internal/domain"
	"TestTask-PGPro/internal/storage/dbModels"
	adapter "TestTask-PGPro/lib/adapter/db"
	"context"
)

type CommandsRepository struct {
	db *adapter.PostgresAdapter
}

func NewCommandsRepository(adapter *adapter.PostgresAdapter) *CommandsRepository {
	return &CommandsRepository{db: adapter}
}

func (c *CommandsRepository) AddCommand(ctx context.Context, text string) (int, error) {
	sql := `insert into commands(text) values ($1) returning id`
	var id int
	err := c.db.ExecuteAndGet(ctx, &id, sql, text)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *CommandsRepository) DeleteCommand(ctx context.Context, id int) error {
	sql := `delete from commands where id = $1`

	return c.db.Execute(ctx, sql, id)
}

func (c *CommandsRepository) GetCommand(ctx context.Context, id int) (domain.Command, error) {
	sql := `select id, text from commands where id = $1`

	var command dbModels.Command
	err := c.db.ExecuteAndGet(ctx, &command, sql, id)

	if err != nil {
		return domain.Command{}, err
	}

	return domain.Command(command), nil
}

func (c *CommandsRepository) GetCommands(ctx context.Context) ([]domain.Command, error) {
	sql := `select id, text from commands`

	var command []dbModels.Command
	err := c.db.Query(ctx, &command, sql)

	if err != nil {
		return nil, err
	}

	return dbModels.DBCommandsToCommands(command), nil
}
