package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCommand())
}

// initCmd represents the init command
func initCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init [name]",
		Short: "Initialize a new jamgo application",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			os.MkdirAll(args[0], os.ModePerm)
			os.Chdir(args[0])
			defer os.Chdir("..")
			os.MkdirAll("pages/templates", os.ModePerm)
			os.MkdirAll("assets/img", os.ModePerm)
			os.MkdirAll("assets/css", os.ModePerm)
			os.MkdirAll("assets/js", os.ModePerm)

			var wg sync.WaitGroup
			wg.Add(5)
			go func() {
				createDefault("pages/templates", "head")
				wg.Done()
			}()
			go func() {
				go createDefault("pages/templates", "header")
				wg.Done()
			}()
			go func() {
				createDefault("pages/templates", "footer")
				wg.Done()
			}()
			go func() {
				cmd := exec.Command("jamgo", "new", "page", "index")
				err := cmd.Run()
				if err != nil {
					log.Fatalln(err)
				}
				wg.Done()
			}()

			go func() {
				cmd := exec.Command("go", "mod", "init", args[0])
				err := cmd.Run()
				if err != nil {
					log.Fatalln(err)
				}
				cmd = exec.Command("go", "get", "github.com/SamtheSaint/jamgo")
				err = cmd.Run()
				if err != nil {
					log.Fatalln(err)
				}
				wg.Done()
			}()

			wg.Wait()
		},
	}
}

func createDefault(path, name string) {
	f, err := os.Create(fmt.Sprintf("%s/%s.gohtml", path, name))
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	r, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/SamtheSaint/jamgo/master/default/%s.gohtml", name))
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()
	io.Copy(f, r.Body)
}
