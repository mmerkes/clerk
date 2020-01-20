package storage

import (
	"encoding/json"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
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

// TODO: Should refactor entire package and separate data layer from modeling layer
func AddTask(task Task) int {
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

	return id
}

func DeleteTask(id int) {
	home, err := homedir.Dir()
	handleError(err)

	db_path := home + "/.clerk-db"

	tasks := loadTasks(db_path)

	for i, t := range tasks.Tasks {
		if t.Id == id {
			tasks.Tasks = remove(tasks.Tasks, i)
			break
		}
	}

	saveTasks(db_path, tasks)
}

func StartTask(id int) {
	home, err := homedir.Dir()
	handleError(err)

	db_path := home + "/.clerk-db"

	tasks := loadTasks(db_path)

	var task *Task

	var index int
	for i, t := range tasks.Tasks {
		if t.Id == id {
			index = i
			task = &t
			break
		}
	}

	if task == nil {
		panic("Task " + string(id) + " does not exist.")
	}

	startTime := time.Now()

	if isTimeUnset(task.StartTime) {
		task.StartTime = startTime
	}

	// TODO: Skip creating event if already started
	event := Event{
		StartTime: startTime,
	}
	task.Events = append(task.Events, event)
	tasks.Tasks[index] = *task

	saveTasks(db_path, tasks)

	isRunning := true
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		isRunning = false
		StopTask(id)
		fmt.Println("")
	}()

	for isRunning {
		printTimeElasped(startTime)
		time.Sleep(1 * time.Second)
	}
}

func printTimeElasped(startTime time.Time) {
	duration := time.Now().Sub(startTime)
	fmt.Printf("\rTime Elapsed: %02.0f:%02.0f:%02.0f", duration.Hours(), duration.Minutes(), duration.Seconds())
}

func StopTask(id int) {
	// TODO: Refactor shared code into function, i.e. getting the DB path, finding a task, etc.
	home, err := homedir.Dir()
	handleError(err)

	db_path := home + "/.clerk-db"

	tasks := loadTasks(db_path)

	var task *Task

	var index int
	for i, t := range tasks.Tasks {
		if t.Id == id {
			index = i
			task = &t
			break
		}
	}

	if task == nil {
		panic("Task " + string(id) + " does not exist.")
	}

	for i, e := range task.Events {
		if isTimeUnset(e.EndTime) {
			e.EndTime = time.Now()
			task.Events[i] = e
		}
	}
	tasks.Tasks[index] = *task

	saveTasks(db_path, tasks)
}

func isTimeUnset(t time.Time) bool {
	emptyTime := time.Time{}
	return emptyTime == t
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
