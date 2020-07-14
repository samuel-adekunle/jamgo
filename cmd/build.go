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
	"sync"

	"go/token"

	"github.com/spf13/cobra"

	"github.com/SamtheSaint/jamgo/tools"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build website in specified directory.",
	Run: func(cmd *cobra.Command, args []string) {
		buildSite()
	},
}

var fset *token.FileSet
var buildDir, rootDir string

func init() {
	rootCmd.AddCommand(buildCmd)
	fset = token.NewFileSet()

	buildCmd.Flags().StringVarP(&buildDir, "dir", "d", "public", "Directory where website is built.")
	rootDir, _ = os.Getwd()
}

func buildSite() {
	var tpl *template.Template
	tpl = template.Must(tpl.ParseGlob("pages/*/*.gohtml"))

	err := os.MkdirAll(buildDir, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	files, err := ioutil.ReadDir("pages")
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(files) - 1)

	for _, f := range files {
		if page := f.Name(); page != "templates" {
			go func() {
				createPageFromTemplate(page, tpl)
				wg.Done()
			}()
		}
	}

	wg.Wait()
}

func getPageData(name string, page chan<- *tools.Page, multiplePage chan<- *[]tools.Page) {
	curDir := fmt.Sprintf("%s/pages/%s", rootDir, name)
	cmd := exec.Command("go", "build", "-buildmode=plugin")
	cmd.Dir = curDir
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	p, err := plugin.Open(fmt.Sprintf("%s/%s.so", curDir, name))
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

	sym, err = p.Lookup("PageDataCollection")
	if err != nil {
		log.Fatalln(err)
	}
	d, _ := sym.(*[]tools.Page)
	multiplePage <- d
	close(multiplePage)
}

func createPageFromTemplate(name string, tpl *template.Template) {
	pageData := make(chan *tools.Page)
	multiplePageData := make(chan *[]tools.Page)
	go getPageData(name, pageData, multiplePageData)

	page, err := os.Create(fmt.Sprintf("%s/%s.html", buildDir, name))
	if err != nil {
		log.Fatalln(err)
	}
	defer page.Close()

	err = tpl.ExecuteTemplate(page, fmt.Sprintf("%s.gohtml", name), <-pageData)
	if err != nil {
		log.Fatalln(err)
	}

	if multiplePage := *<-multiplePageData; multiplePage != nil {
		var wg sync.WaitGroup
		wg.Add(len(multiplePage))
		err := os.MkdirAll(fmt.Sprintf("%s/%s", buildDir, name), os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		for _, pageData := range multiplePage {
			go func(p tools.Page) {
				page, err := os.Create(fmt.Sprintf("%s/%s/%s | %s.html", buildDir, name, name, p.Title))
				if err != nil {
					log.Fatalln(err)
				}
				defer page.Close()

				err = tpl.ExecuteTemplate(page, fmt.Sprintf("%s.gohtml", name), p)
				if err != nil {
					log.Fatalln(err)
				}
				wg.Done()
			}(pageData)
		}
		wg.Wait()
	}
}
