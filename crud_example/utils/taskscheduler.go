package utils

import (
	"crud_example/initializers"
	"crud_example/models"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func Executor() {
	fmt.Println("Start of the task processing...")
	var ids []int
	initializers.DB.Where("status = ? ", os.Getenv("INPSTATUS")).Find(&ids)
	for _, id := range ids {
		fmt.Println("Processing the Task ID = ", id)
		processTask(id)
	}
}

func processTask(id int) {
	initializers.DB.Model(&models.Task{}).Where("id = ?", id).Update("status", os.Getenv("INPSTATUS"))
	if err := initializers.DB.Model(&models.Task{}).Where("id = ?", id).Update("output", gorm.Expr("upper(input)")).Error; err != nil {
		fmt.Println("Error while processing the task, id = ", id)
		initializers.DB.Model(&models.Task{}).Where("id = ?", id).Update("status", os.Getenv("ERRSTATUS"))
		return
	}
	initializers.DB.Model(&models.Task{}).Where("id = ?", id).Update("status", os.Getenv("CMPSTATUS"))
}
