package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	newCmd.AddCommand(templateCommand())
}

// templateCommand represents the template command
func templateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "template [name]",
		Short: "Create a new template, in the pages/templates directory.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			defer func() {
				if r := recover(); r != nil {
					log.Println("Try executing command in root directory of application.")
				}
			}()
			f, err := os.Create(fmt.Sprintf("pages/templates/%s.gohtml", args[0]))
			if err != nil {
				log.Panicln(err)
			}
			defer f.Close()
			writeTemplate(f, args[0])
		},
	}
}

func writeTemplate(f *os.File, name string) {
	f.WriteString(fmt.Sprintf("{{template \"%s\"}}\n", name))
	f.WriteString("{{end}}")
}
