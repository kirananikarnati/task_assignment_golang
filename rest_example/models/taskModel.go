package models

type Task struct {
	ID       int    `json:"id"`
	Input    string `json:"input"`
	Status   string `json:"status"`
	Output   string `json:"output"`
	ErrorMsg string `json:"error"`
}
