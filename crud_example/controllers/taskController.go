package controllers

import (
	"crud_example/initializers"
	"crud_example/models"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var reqBody struct {
		Input string
	}
	c.Bind(&reqBody)

	task := models.Task{Input: reqBody.Input, Status: os.Getenv("TaskInitialStatus")}
	result := initializers.DB.Create(&task)
	if result.Error != nil {
		log.Fatal("Failed to create the task.")
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"id":  task.ID,
		"msg": "task created",
	})
}

func RetrieveAllTasks(c *gin.Context) {
	var tasks []models.Task
	initializers.DB.Find(&tasks)
	c.JSON(200, gin.H{
		"tasksList": tasks,
	})
}

func RetrieveTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	initializers.DB.First(&task, id)
	c.JSON(200, gin.H{
		"task": task,
	})
}

func TaskStatus(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	initializers.DB.First(&task, id)
	c.JSON(200, gin.H{
		"task_status": task.Status,
	})
}

func TaskResult(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	initializers.DB.First(&task, id)
	c.JSON(200, gin.H{
		"task_result": task.Output,
	})
}

func ProcessTask(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("in process task...")
	var task models.Task
	initializers.DB.First(&task, id)
	if strings.ToUpper(task.Input) == task.Output {
		c.JSON(200, gin.H{
			"result": "Task is already processed.",
		})
	} else {
		initializers.DB.Model(&task).Where("id = ?", id).Updates(models.Task{Status: os.Getenv("CMPSTATUS"), Output: strings.ToUpper(task.Input)})
		c.JSON(200, gin.H{
			"result": "Task processing is successful.",
		})
	}
}

//.Update("output", gorm.Expr("upper(input)"))
