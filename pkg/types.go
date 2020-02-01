package clerk

import "time"

type TaskStatus string

const (
	Created  TaskStatus = "Created"
	Started  TaskStatus = "Started"
	Complete TaskStatus = "Complete"
)

func IsValidTaskStatus(status TaskStatus) bool {
	switch status {
	case Created, Started, Complete:
		return true
	}

	return false
}

type Event struct {
	StartTime time.Time
	EndTime   time.Time
}

type Task struct {
	Id          int
	Title       string
	Description string
	Events      []Event
	CreateTime  time.Time
	StartTime   time.Time
	EndTime     time.Time
}

type Tasks struct {
	Tasks []Task
}
