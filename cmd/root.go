package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var defaultPath string
var path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gazo",
	Short: "Create a temporary PNG image.",
	Long:  "Create a temporary PNG image. \nUse \"gazo img --help\" for more information about a command.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defaultPath = home + string(os.PathSeparator)

	rootCmd.PersistentFlags().StringVar(&path, "path", defaultPath, "Output Directory")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
