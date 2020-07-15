package cmd

import (
	"fmt"
	"os"
	"testing"
)

func Test_createDefault(t *testing.T) {
	type args struct {
		path string
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"default creation test",
			args{
				"testdata",
				"page",
			},
		},
		{
			"secondary test",
			args{
				"testdata",
				"header",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createDefault(tt.args.path, tt.args.name)
			_, err := os.Open(fmt.Sprintf("%s/%s.gohtml", tt.args.path, tt.args.name))
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func Benchmark_createDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createDefault("testdata", "page")
	}
}

func Example_initCommand() {
	os.Chdir("testdata")
	defer os.Chdir("..")
	cmd := initCommand()
	cmd.SetArgs([]string{"tmp"})
	cmd.Execute()
	//Output:
}
