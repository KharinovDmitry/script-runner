package adapter

import (
	"context"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

type PostgresAdapter struct {
	connection *sqlx.DB
	TimeoutDb  int
}

func NewPostgresAdapter(timeoutDbInSecond int) *PostgresAdapter {
	return &PostgresAdapter{
		TimeoutDb: timeoutDbInSecond,
	}
}

func (p *PostgresAdapter) Connect(ctx context.Context, connectionString string) (*sqlx.DB, error) {
	conn, err := sqlx.ConnectContext(ctx, "postgres", connectionString)
	p.connection = conn
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (p *PostgresAdapter) Close() error {
	return p.connection.Close()
}

func (p *PostgresAdapter) Execute(requestCtx context.Context, sql string, args ...interface{}) error {
	if p.connection == nil {
		return errors.New("[ PostgresAdapter ] Execute: connection is nil")
	}

	_, cancel := context.WithTimeout(requestCtx, time.Duration(p.TimeoutDb)*time.Second)
	defer cancel()

	_, err := p.connection.Exec(sql, args...)
	return err
}

func (p *PostgresAdapter) ExecuteAndGet(requestCtx context.Context, destination interface{}, sql string, args ...interface{}) error {
	if p.connection == nil {
		return errors.New("[ PostgresAdapter ] ExecuteAndGet: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, time.Duration(p.TimeoutDb)*time.Second)
	defer cancel()

	return p.connection.GetContext(ctx, destination, sql, args...)
}

func (p *PostgresAdapter) Query(requestCtx context.Context, destination interface{}, query string, args ...interface{}) error {
	if p.connection == nil {
		return errors.New("[ PostgresAdapter ] Query: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, time.Duration(p.TimeoutDb)*time.Second)
	defer cancel()

	return p.connection.SelectContext(ctx, destination, query, args...)
}
