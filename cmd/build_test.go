package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

func Example_buildCommand() {
	viper.SetConfigFile("../.jamgo")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	dir := viper.GetString("DIR")

	os.Chdir("testdata")
	defer os.Chdir("..")

	cmd := initCommand()
	cmd.SetArgs([]string{"buildTest"})
	cmd.Execute()

	os.Chdir("buildTest")
	defer os.Chdir("..")
	c := exec.Command("go", "mod", "edit", "-replace", fmt.Sprintf("github.com/SamtheSaint/jamgo=%s", dir))
	if e := c.Run(); e != nil {
		log.Fatalln(e)
	}

	buildCommand().Execute()
	//Output:
}
