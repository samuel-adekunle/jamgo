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
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Initialize a new jamgo application",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(args[0], os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		err = os.Chdir(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		var wg sync.WaitGroup

		os.MkdirAll("pages/templates", os.ModePerm)
		os.MkdirAll("assets/img", os.ModePerm)
		os.MkdirAll("assets/css", os.ModePerm)
		os.MkdirAll("assets/js", os.ModePerm)

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
			cmd = exec.Command("go", "mod", "edit", "-require", "github.com/SamtheSaint/jamgo@master")
			err = cmd.Run()
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()

		wg.Wait()
	},
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

func init() {
	rootCmd.AddCommand(initCmd)
}
