package mapper

import (
	"esdc-backend/internal/module/admin/dto"
	"esdc-backend/internal/module/common/model"
	"time"
)

func MapToUserDataForAdmin(users *[]model.User) *[]dto.UserDataForAdmin {
	if users == nil {
		return nil
	}
	var filteredUsers []dto.UserDataForAdmin
	for _, user := range *users {
		filteredUsers = append(filteredUsers, dto.UserDataForAdmin{
			ID:             user.ID,
			Name:           user.Name,
			Email:          user.Email,
			Username:       user.Username,
			GithubUsername: user.Github.Username,
			Role:           user.Role,
			Status:         user.Status,
			CreatedAt:      getCreatedDateFromNumber(user.CreatedAt),
			UpdatedAt:      getCreatedDateFromNumber(user.UpdatedAt),
		})
	}
	return &filteredUsers
}
func getCreatedDateFromNumber(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
