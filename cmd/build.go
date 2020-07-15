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

var fset *token.FileSet
var buildDir, rootDir string

func init() {
	cmd := buildCommand()
	fset = token.NewFileSet()
	cmd.Flags().StringVarP(&buildDir, "dir", "d", "public", "Directory where website is built.")
	rootCmd.AddCommand(cmd)
}

// buildCommand represents the build command
func buildCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Build website in specified directory.",
		Run: func(cmd *cobra.Command, args []string) {
			var tpl *template.Template
			rootDir, _ = os.Getwd()
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
		},
	}

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

	if p := <-pageData; p != nil {
		page, err := os.Create(fmt.Sprintf("%s/%s.html", buildDir, name))
		if err != nil {
			log.Fatalln(err)
		}
		defer page.Close()

		err = tpl.ExecuteTemplate(page, fmt.Sprintf("%s.gohtml", name), p)
		if err != nil {
			log.Fatalln(err)
		}
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

				err = tpl.ExecuteTemplate(page, fmt.Sprintf("%s_multiple.gohtml", name), p)
				if err != nil {
					log.Fatalln(err)
				}
				wg.Done()
			}(pageData)
		}
		wg.Wait()
	}
}
