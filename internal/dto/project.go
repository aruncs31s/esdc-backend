package dto

import "time"

// ProjectCreation represents project creation request
// @Description Project creation request payload
type ProjectCreation struct {
	Title        string    `json:"title" example:"My Project"`                                 // Project title
	Image        *string   `json:"image" example:"https://example.com/image.jpg"`              // Project image URL
	Description  string    `json:"description" example:"This is a sample project description"` // Project description
	GithubLink   string    `json:"github_link" example:"https://github.com/user/project"`      // Project link
	Tags         *[]string `json:"tags" example:"go,api,backend"`                              // Project tags
	Contributers *[]string `json:"contributers" example:"user1,user2,user3"`                   // Contributor user IDs
	Technologies *[]string `json:"technologies" example:"Go, Gin, GORM"`                       // Technologies used
	LiveUrl      *string   `json:"live_url" example:"https://example.com/live"`                // Live URL of the project
}
type ProjectResponse struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Technologies *[]string  `json:"technologies"`
	Image        *string    `json:"image"`
	GithubLink   string     `json:"github_link"`
	LiveUrl      *string    `json:"live_url"`
	Contributors []string   `json:"contributors"`
	Tags         []string   `json:"tags"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	// Newly Addedd
	Likes    int    `json:"likes"`
	Cost     int    `json:"cost"`
	Category string `json:"category"`
}
