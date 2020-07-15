package cmd

import (
	"log"
	"os"
)

func Example_pageCommand() {
	err := os.Chdir("testdata")
	if err != nil {
		log.Println(os.Getwd())
		log.Fatalln(err, "failed here")
	}
	defer os.Chdir("..")
	os.MkdirAll("pages", os.ModePerm)
	cmd := pageCommand()
	cmd.SetArgs([]string{"index"})
	cmd.Execute()
	//Output:
}
