/*
Copyright © 2020 Samuel Adekunle <ebunsamuel@yahoo.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
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

	f.WriteString("{{template \"head\" .Title}}\n")
	f.WriteString("{{template \"header\"}}\n")
	f.WriteString("<main></main>\n")
	f.WriteString("{{template \"footer\"}}\n")

	g, err := os.Create(fmt.Sprintf("pages/%s/%s.go", name, name))
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()
	starterString := `package main

import "github.com/SamtheSaint/jamgo/tools"

// PageData supples data for the page to parse
var PageData tools.Page

func init() {
	PageData = tools.Page{
		Title: "Index",
		Data:  nil,
	}
}`
	g.WriteString(starterString)

}
