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
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build website in public directory",
	// REVIEW - add Long description
	Long: `Longer description here.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildSite()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func buildSite() {
	err := os.MkdirAll("public", os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	var tpl *template.Template
	tpl = template.Must(tpl.ParseGlob("pages/*/*.gohtml"))
	files, err := ioutil.ReadDir("pages")
	if err != nil {
		log.Fatalln(err)
	}
	wg.Add(len(files) - 1)
	for _, f := range files {
		if page := f.Name(); page != "templates" {
			fmt.Println(page)
			wg.Done()
		}
	}
	wg.Wait()
}

func createPageFromTemplate(name string, tpl *template.Template) {

}
