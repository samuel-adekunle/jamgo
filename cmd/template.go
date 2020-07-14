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
	"log"
	"os"

	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template [name]",
	Short: "Create a new template, in the template directory.",
	// REVIEW - add Long description
	Long: `Longer description here.`,
	Args: cobra.ExactArgs(1),
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
