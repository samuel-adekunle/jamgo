package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new resource.",
}

func init() {
	rootCmd.AddCommand(newCmd)
}
