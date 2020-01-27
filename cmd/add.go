/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/mmerkes/clerk/pkg"
	"github.com/spf13/cobra"

	"bufio"
	"fmt"
	"os"
)

// TODO: Rename *-task command to remove task, which can be assumed
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task",
	Long:  `Adds a new task to your task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Try to set the root persistence directory.
		clerk.SetRootPersistenceDir(rootPersistenceDir)
		task := clerk.Task{}

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Enter Title:")
		scanner.Scan()
		task.Title = scanner.Text()

		fmt.Println("Enter Description:")
		scanner.Scan()
		task.Description = scanner.Text()

		task_id := clerk.AddTask(task)
		fmt.Printf("Added task with ID %d\n", task_id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
