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

func init() {
	newCmd.AddCommand(pageCmd)
}

func createPage(name string) {
	err := os.MkdirAll(fmt.Sprintf("pages/%s", name), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	wg.Add(2)
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
