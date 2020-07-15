package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template [name]",
	Short: "Create a new template, in the pages/templates directory.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createTemplate(args[0])
	},
}

func init() {
	newCmd.AddCommand(templateCmd)
}

func createTemplate(name string) {
	f, err := os.Create(fmt.Sprintf("pages/templates/%s.gohtml", name))
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("{{template \"%s\"}}\n", name))
	f.WriteString("{{end}}")
}
