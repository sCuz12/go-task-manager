package task

import (
	"testing"
	"time"
)



func TestAddTask(t *testing.T) {
	tm := NewTaskManager()

	taskID := tm.AddTask("Test Task",time.Now().Add(2 * time.Second),func(){})

	if _, exists := tm.Tasks[taskID]; !exists {
        t.Errorf("Expected task with ID %v to be added", taskID)
    }

}


func TestRemoveTask(t *testing.T) {
	tm:=NewTaskManager()

	taskID := tm.AddTask("Test Task",time.Now().Add(2 * time.Second),func() {})

	delete(tm.Tasks,taskID)

	if _, exists := tm.Tasks[taskID]; exists {
        t.Errorf("Expected task with ID %v to be removed", taskID)
    }
}

// Test task execution
func TestTaskExecution(t *testing.T) {
    tm := NewTaskManager()
    executed := false
    tm.AddTask("Test Task", time.Now().Add(1*time.Second), func() {
        executed = true
    })
    
    go tm.Start()
    time.Sleep(4 * time.Second) // Wait to allow the task to be executed
    // tm.Stop()

    if !executed {
        t.Errorf("Expected task to be executed")
    }
}