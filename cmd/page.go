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

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page [path]",
	Short: "Create a new page, relative to pages directory.",
	// REVIEW - add Long description
	Long: `Longer description here.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createPage(args[0])
	},
}

func init() {
	newCmd.AddCommand(pageCmd)
}

func createPage(name string) {
	err := os.MkdirAll(fmt.Sprintf("pages/%s", name), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(fmt.Sprintf("pages/%s/%s.gohtml", name, name))
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	f.WriteString("{{template \"head\"}}\n")
	f.WriteString("{{template \"header\"}}\n")
	f.WriteString("<main></main>\n")
	f.WriteString("{{template \"footer\"}}\n")
}
