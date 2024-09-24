/*
Copyright © 2024 ktsnkmrcom
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gazo v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
