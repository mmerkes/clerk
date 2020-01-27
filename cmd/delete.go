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
)

var id int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Delete a task from your set of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Try to set the root persistence directory.
		clerk.SetRootPersistenceDir(rootPersistenceDir)
		clerk.DeleteTask(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().IntVarP(&id, "id", "i", -1, "Id of task to delete")
	deleteCmd.MarkFlagRequired("id")
}
