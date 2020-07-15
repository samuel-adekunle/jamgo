package cmd

import (
	"os"
)

func Example() {
	os.Chdir("testdata")
	cmd := initCommand()
	cmd.SetArgs([]string{"buildTest"})
	cmd.Execute()

	os.Chdir("buildTest")
	cmd = buildCommand()
	cmd.Execute()
	//Output:
}
