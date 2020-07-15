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

var multiple bool

func init() {
	cmd := pageCommand()
	cmd.Flags().BoolVarP(&multiple, "multiple", "m", false, "Toggle creation of multi-page template")
	newCmd.AddCommand(cmd)
}

func pageCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "page [name]",
		Short: "Create a new page, in pages/name directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var wg sync.WaitGroup
			wg.Add(2)
			os.MkdirAll(fmt.Sprintf("pages/%s", args[0]), os.ModePerm)

			go func() {
				createPage(args[0], "gohtml")
				wg.Done()
			}()

			go func() {
				createPage(args[0], "go")
				wg.Done()
			}()

			if multiple {
				wg.Add(1)
				go func() {
					createPage(fmt.Sprintf("%s_multiple", args[0]), "gohtml")
					wg.Done()
				}()
			}
			wg.Wait()
		},
	}
}

func createPage(name, extension string) {
	f, err := os.Create(fmt.Sprintf("pages/%s/%s.%s", name, name, extension))
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	r, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/SamtheSaint/jamgo/master/default/page.%s", extension))
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	io.Copy(f, r.Body)
}
