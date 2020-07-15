package cmd

import (
	"io/ioutil"
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"Execute Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetOut(ioutil.Discard)
			Execute()
		})
	}
}
