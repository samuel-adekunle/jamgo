package cmd

import (
	"os"
)

func Example_pageCommand() {
	os.Chdir("testdata")
	os.MkdirAll("pages", os.ModePerm)
	cmd := pageCommand()
	cmd.SetArgs([]string{"index"})
	cmd.Execute()
	os.Chdir("..")
	//Output:
}
