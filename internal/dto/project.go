package dto

// ProjectCreation represents project creation request
// @Description Project creation request payload
type ProjectCreation struct {
	Name         string   `json:"name" example:"My Project"`                                    // Project name
	Image        string   `json:"image" example:"https://example.com/image.jpg"`               // Project image URL
	Description  string   `json:"description" example:"This is a sample project description"`   // Project description
	Link         string   `json:"link" example:"https://github.com/user/project"`               // Project link
	Tags         []string `json:"tags" example:"go,api,backend"`                               // Project tags
	Contributers []int    `json:"contributers" example:"1,2,3"`                               // Contributor user IDs
}


