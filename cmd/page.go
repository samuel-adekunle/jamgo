package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page [name]",
	Short: "Create a new page, in pages/name directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createPage(args[0])
	},
}

var multiple bool

func init() {
	newCmd.AddCommand(pageCmd)

	pageCmd.Flags().BoolVarP(&multiple, "multiple", "m", false, "Toggle creation of multi-page template")
}

func createPage(name string) {
	var wg sync.WaitGroup
	if multiple {
		wg.Add(3)
	} else {
		wg.Add(2)
	}

	err := os.MkdirAll(fmt.Sprintf("pages/%s", name), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		f, err := os.Create(fmt.Sprintf("pages/%s/%s.gohtml", name, name))
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		r, err := http.Get("https://raw.githubusercontent.com/SamtheSaint/jamgo/master/default/page.gohtml")
		if err != nil {
			log.Fatalln(err)
		}
		defer r.Body.Close()

		io.Copy(f, r.Body)
		wg.Done()
	}()

	if multiple {
		go func() {
			f, err := os.Create(fmt.Sprintf("pages/%s/%s_multiple.gohtml", name, name))
			if err != nil {
				log.Fatalln(err)
			}
			defer f.Close()

			r, err := http.Get("https://raw.githubusercontent.com/SamtheSaint/jamgo/master/default/page.gohtml")
			if err != nil {
				log.Fatalln(err)
			}
			defer r.Body.Close()

			io.Copy(f, r.Body)
			wg.Done()
		}()
	}

	go func() {
		g, err := os.Create(fmt.Sprintf("pages/%s/%s.go", name, name))
		if err != nil {
			log.Fatalln(err)
		}
		defer g.Close()

		p, err := http.Get("https://raw.githubusercontent.com/SamtheSaint/jamgo/master/default/page.go")
		if err != nil {
			log.Fatalln(err)
		}
		defer p.Body.Close()

		io.Copy(g, p.Body)
		wg.Done()
	}()

	wg.Wait()
}
