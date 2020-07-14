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
	"os/exec"
	"plugin"

	"go/token"

	"github.com/spf13/cobra"

	"github.com/SamtheSaint/jamgo/tools"
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

var fset *token.FileSet

func init() {
	rootCmd.AddCommand(buildCmd)
	fset = token.NewFileSet()
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
			go createPageFromTemplate(page, tpl)
		}
	}
	wg.Wait()
}

func getPageData(name string, page chan<- *tools.Page) {
	os.Chdir(fmt.Sprintf("pages/%s", name))
	cmd := exec.Command("go", "build", "-buildmode=plugin")
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	p, err := plugin.Open(fmt.Sprintf("%s.so", name))
	if err != nil {
		log.Fatalln(err)
	}
	sym, err := p.Lookup("PageData")
	if err != nil {
		log.Fatalln(err)
	}
	data, _ := sym.(*tools.Page)
	page <- data
	close(page)
}

func createPageFromTemplate(name string, tpl *template.Template) {
	pageData := make(chan *tools.Page)
	go getPageData(name, pageData)
	page, err := os.Create(fmt.Sprintf("public/%s.html", name))
	if err != nil {
		log.Fatalln(err)
	}
	defer page.Close()
	err = tpl.ExecuteTemplate(page, fmt.Sprintf("%s.gohtml", name), <-pageData)
	if err != nil {
		log.Fatalln(err)
	}
	wg.Done()
}
