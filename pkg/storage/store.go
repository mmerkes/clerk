package storage

import (
	"encoding/json"
	homedir "github.com/mitchellh/go-homedir"
	"io/ioutil"
	"strings"
	"time"
)

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

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func setDefaultTaskValues(task *Task) {
	task.Events = []Event{}
	task.CreateTime = time.Now()
}

func AddTask(task Task) {
	setDefaultTaskValues(&task)

	home, err := homedir.Dir()
	handleError(err)

	db_path := home + "/.clerk-db"

	tasks := loadTasks(db_path)

	id := 1

	for _, t := range tasks.Tasks {
		if t.Id >= id {
			id = t.Id + 1
		}
	}
	task.Id = id

	tasks.Tasks = append(tasks.Tasks, task)

	saveTasks(db_path, tasks)
}

func DeleteTask(id int) {
	home, err := homedir.Dir()
	handleError(err)

	db_path := home + "/.clerk-db"

	tasks := loadTasks(db_path)

	for i, t := range tasks.Tasks {
		if t.Id == id {
			tasks.Tasks = remove(tasks.Tasks, i)
		}
	}

	saveTasks(db_path, tasks)
}

func remove(slice []Task, i int) []Task {
	return append(slice[:i], slice[i+1:]...)
}

func loadTasks(db_path string) Tasks {
	tasks := Tasks{}
	raw_tasks, err := ioutil.ReadFile(db_path)
	if err != nil {
		if strings.HasSuffix(err.Error(), "no such file or directory") {
			tasks.Tasks = []Task{}
			return tasks
		}
		handleError(err)
	}

	err = json.Unmarshal([]byte(raw_tasks), &tasks)
	handleError(err)

	return tasks
}

func saveTasks(db_path string, tasks Tasks) {
	s, err := json.MarshalIndent(tasks, "", "  ")
	handleError(err)

	err = ioutil.WriteFile(db_path, s, 0644)
	handleError(err)
}
