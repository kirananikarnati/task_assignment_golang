package main

import (
	"crud_example/controllers"
	"crud_example/initializers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/createtask", controllers.CreateTask)
	r.GET("/taskstatus/:id", controllers.RetrieveTask)
	r.GET("/tastresult/:id", controllers.RetrieveTask)

	r.GET("/retrievealltasks", controllers.RetrieveAllTasks)
	r.GET("/retrievetask/:id", controllers.RetrieveTask)
	r.PUT("/processtask/:id", controllers.ProcessTask)

	// s := gocron.NewScheduler()
	// s.Every(3).Seconds().Do(utils.Executor)
	// go s.Start()
	r.Run()

}

// func Executor() {
// 	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
// 	if err != nil {
// 		fmt.Println("error connecting to db")
// 		panic(err)
// 	}

// 	var wg sync.WaitGroup
// 	es := make(chan func(), 10)
// 	for i := 0; i < 1; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			es <- func() {
// 				var id int
// 				var input string
// 				err := db.QueryRow("SELECT id, status FROM tasks WHERE status = $1", os.Getenv("TaskInitialStatus")).Scan(&id, &input)
// 				if err != nil {
// 					fmt.Println(err)
// 					return
// 				}

// 				fmt.Printf("Processing record %d \n", id)

// 				// time.Sleep(time.Second)

// 				_, err = db.Exec("UPDATE tasks SET status = $1 WHERE id = $2", os.Getenv("INPSTATUS"), id)
// 				if err != nil {
// 					fmt.Println(err)
// 					return
// 				}

// 				_, err = db.Exec("UPDATE tasks SET output = upper(input) WHERE id = $1", id)
// 				if err != nil {
// 					fmt.Println(err)
// 					db.Exec("UPDATE tasks SET status = $1 WHERE id = $2", os.Getenv("ERRSTATUS"), id)
// 					return
// 				}

// 				_, err = db.Exec("UPDATE tasks SET status = $1 WHERE id = $2", os.Getenv("CMPSTATUS"), id)
// 				if err != nil {
// 					fmt.Println(err)
// 					return
// 				}
// 				fmt.Printf("Completed processing record %d\n", id)
// 			}
// 		}(i)
// 	}

// 	go func() {
// 		for task := range es {
// 			task()
// 		}
// 	}()

// 	wg.Wait()
// 	close(es)
// 	// defer db.Close()
// }
