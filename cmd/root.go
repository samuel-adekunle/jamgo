package cmd

import (
	"log"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
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

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
	}

	// Search config in home directory with name ".jamgo" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".jamgo")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}
