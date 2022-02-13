package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/open-blockchain-explorer/tnbassist/testdata"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func Test_OverviewCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
		regex   bool
	}{
		{
			name:    "Should throw error if necessary argument \"file\" is not provided",
			args:    []string{},
			want:    fmt.Sprintf("Error: %s\n%s\n", testdata.RequiredFlagMsg("file"), testdata.StatsOverviewUsage),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			rootCmd.SetOut(b)
			rootCmd.SetErr(b)
			rootCmd.SetArgs(append([]string{statsCmd.Name(), overviewCmd.Name()}, tt.args...))

			if _, err := rootCmd.ExecuteC(); (err != nil) != tt.wantErr {
				t.Errorf("%sCmd.Execute() error = %v, wantErr %v", overviewCmd.Name(), err, tt.wantErr)
			}
			backupCmd.Flags().VisitAll(func(flag *pflag.Flag) {
				flag.Value.Set(flag.DefValue)
			})
			out, err := ioutil.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			if tt.regex {
				assert.Regexp(t, tt.want, string(out))
			} else {
				assert.Equal(t, tt.want, string(out))
			}
		})
	}
}
