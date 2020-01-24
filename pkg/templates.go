package clerk

import "io/ioutil"

func shortTasksTemplate() ([]byte, error) {
	return ioutil.ReadFile("templates/tasks-short.tmpl")
}

func verboseTasksTemplate() ([]byte, error) {
	return ioutil.ReadFile("templates/tasks-verbose.tmpl")
}
