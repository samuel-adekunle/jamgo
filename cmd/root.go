package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jamgo",
	Short: "Jamgo is a minimal, superfast golang static site generator",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
	//ANCHOR - uncomment to generate docs
	// if err := doc.GenMarkdownTree(rootCmd, "./docs"); err != nil {
	// 	log.Fatalln(err)
	// }
}
