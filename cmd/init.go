/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

var wg sync.WaitGroup

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new jamgo application",
	// REVIEW - add Long description
	Long: `Longer description here.`,

	Run: func(cmd *cobra.Command, args []string) {
		os.MkdirAll("pages/templates", os.ModePerm)
		os.MkdirAll("assets/img", os.ModePerm)
		os.MkdirAll("assets/css", os.ModePerm)
		os.MkdirAll("assets/js", os.ModePerm)

		wg.Add(4)
		go createDefault("pages", "index")
		go createDefault("pages/templates", "head")
		go createDefault("pages/templates", "header")
		go createDefault("pages/templates", "footer")
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
	wg.Done()
}

func init() {
	rootCmd.AddCommand(initCmd)
}
