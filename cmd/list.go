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
	"fmt"
	"strings"

	"github.com/mmerkes/clerk/pkg"
	"github.com/spf13/cobra"
)

var Verbose bool
var Statuses string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks in the console",
	Long:  `Prints a list of tasks in the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Try to set the root persistence directory.
		clerk.SetRootPersistenceDir(rootPersistenceDir)

		statuses := []string{}
		if Statuses != "" {
			statuses = strings.Split(Statuses, ",")
		}

		filterStatuses := make([]clerk.TaskStatus, len(statuses))
		for i, status := range statuses {
			filterStatuses[i] = clerk.TaskStatus(status)
			if !clerk.IsValidTaskStatus(filterStatuses[i]) {
				panic(fmt.Sprintf("%s is not a valid status", status))
			}
		}

		clerk.ListTasks(Verbose, filterStatuses)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	listCmd.PersistentFlags().StringVar(&Statuses, "status", "",
		"Task status filter to restrict statuses or see completed tasks. "+
			"Can be Created, Started, Complete or a comma-separated combination")
}
