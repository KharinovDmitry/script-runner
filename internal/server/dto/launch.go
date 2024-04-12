package dto

import "TestTask-PGPro/internal/domain"

type Launch struct {
	ID        int    `json:"id"`
	CommandID int    `json:"commandID"`
	Output    string `json:"output"`
}

func LaunchesToLaunchesDTO(launches []domain.Launch) []Launch {
	res := make([]Launch, len(launches))
	for i, _ := range res {
		res[i] = Launch((launches[i]))
	}
	return res
}
