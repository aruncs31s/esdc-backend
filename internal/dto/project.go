package dto

type ProjectCreation struct {
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Description  string   `json:"description"`
	Link         string   `json:"link"`
	Tags         []string `json:"tags"`
	Contributers []int    `json:"contributers"`
}


