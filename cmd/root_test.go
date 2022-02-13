package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/open-blockchain-explorer/tnbassist/testdata"
	"github.com/stretchr/testify/assert"
)

func Test_RootCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name:    "Should return help when no args are given to root command",
			args:    []string{},
			want:    testdata.HelpText,
			wantErr: false,
		},
		{
			name:    "Should return version informarion when version flag is given to root command",
			args:    []string{"-v"},
			want:    fmt.Sprintf("%s version %s\n", rootCmd.Name(), rootCmd.Version),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			rootCmd.SetOut(b)
			rootCmd.SetErr(b)
			rootCmd.SetArgs(tt.args)
			if err := rootCmd.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("rootCmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			out, err := ioutil.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, string(out))
		})
	}
}
