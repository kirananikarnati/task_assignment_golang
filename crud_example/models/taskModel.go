package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Input    string `json:"input"`
	Status   string `json:"status"`
	Output   string `json:"output"`
	ErrorMsg string `json:"error_msg"`
}
