package main

import (
	"fmt"
	"taskScheduler/task"
	"time"
)

func main() {


	tm := task.NewTaskManager()
	testTasks := 50


	for i := 0; i < testTasks; i++ {
		taskIndex := i // Capture the current value of i
		tm.AddTask(
			fmt.Sprintf("Task%v", taskIndex),
			time.Now().Add(10 * time.Second),
			func() {
				fmt.Println(fmt.Sprintf("Task%v IS executing", taskIndex))
			},
		)
	}
	
	tm.Start()
}