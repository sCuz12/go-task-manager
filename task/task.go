package task

import (
	"time"

	"github.com/google/uuid"
)


type Task struct {
	ID uuid.UUID
	Name string
	Schedule time.Time
	Action func()
}
