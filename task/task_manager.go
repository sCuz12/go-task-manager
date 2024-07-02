package task

import (
	"math/big"
	"sync"
	"time"

	"github.com/google/uuid"
)

type TaskManager struct {
	Tasks map[uuid.UUID] *Task // map with key[id] for each Task
	IsRunning bool
	mu sync.Mutex //protect access to Task map from different go routines
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make(map[uuid.UUID] *Task),
		IsRunning: false,
	}
}


func (tm *TaskManager) AddTask(name string ,schedule time.Time , action func()) uuid.UUID{
	//validate task
	tm.mu.Lock()
	defer tm.mu.Unlock() //unlock on end of add task
	taskID := generateUUIDInt()


	task := Task{
		ID: taskID,
		Name: name,
		Schedule: schedule,
		Action: action,

	}

	tm.Tasks[taskID] = &task

	return taskID
}

func (tm *TaskManager) Start() {
	tm.IsRunning = true
	
	for {

		if !tm.IsRunning {
			break
		}

		tm.checkAndRunTasks()

		time.Sleep(1 * time.Second)
	}
}

func (tm *TaskManager) checkAndRunTasks() {
	tm.mu.Lock()
	for _,task := range tm.Tasks {
		if task.Schedule.Before(time.Now()){
			go task.Action()	
			delete(tm.Tasks,task.ID)
		}
	 }
	 tm.mu.Unlock()
	 time.Sleep(1 * time.Second)
}


func generateUUIDInt()  uuid.UUID{
    u := uuid.New()
    i := new(big.Int)
    i.SetString(u.String(), 16)
    return u
}