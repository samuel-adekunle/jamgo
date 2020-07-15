package cmd

import (
	"os"
	"testing"
)

func Example_writeTemplate() {
	writeTemplate(os.Stdout, "componentX")

	// Output:
	// {{template "componentX"}}
	// {{end}}
}

func Benchmark_writeTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		writeTemplate(nil, "componentX")
	}
}

func Example_templateCommand() {
	cmd := templateCommand()
	cmd.SetArgs([]string{"componentX"})
	cmd.Execute()
	// Output:
}
