package dbModels

import "TestTask-PGPro/internal/domain"

type Launch struct {
	ID        int    `db:"id"`
	CommandID int    `db:"command_id"`
	Output    string `db:"output"`
}

func DBLaunchesToLaunches(launches []Launch) []domain.Launch {
	res := make([]domain.Launch, len(launches))
	for i, _ := range res {
		res[i] = domain.Launch(launches[i])
	}
	return res
}
