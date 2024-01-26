package models

type StudentCard struct {
	Error       int    `json:"error"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	Path        string `json:"path"`
	PathDoctype string `json:"path_doc"`
}
