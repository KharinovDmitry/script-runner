package storage

import (
	domain "TestTask-PGPro/internal/domain/repository"
	"TestTask-PGPro/internal/storage/postgres"
	adapter "TestTask-PGPro/lib/adapter/db"
)

type Storage struct {
	DB *adapter.PostgresAdapter

	CommandsRepository domain.ICommandsRepository
	LaunchesRepository domain.ILaunchesRepository
}

func NewStorage(db *adapter.PostgresAdapter) *Storage {
	return &Storage{
		DB: db,
	}
}

func (s *Storage) InitStorage() {
	s.CommandsRepository = postgres.NewCommandsRepository(s.DB)
	s.LaunchesRepository = postgres.NewLaunchesRepository(s.DB)
}
