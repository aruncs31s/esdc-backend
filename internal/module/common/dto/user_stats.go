package dto

type UsersStats struct {
	TotalUsers      int `json:"total_users"`
	TotalProjects   int `json:"total_projects"`
	TotalChallenges int `json:"total_challenges"`
	ActiveUsers     int `json:"active_users"`
}
