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

func Test_BackupCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
		regex   bool
	}{
		// Accounts
		{
			name:    "Should backup accounts",
			args:    []string{"accounts", "-H", "127.0.0.1"},
			want:    testdata.FakeAccountsJSON,
			wantErr: false,
		},
		// Wrong command
		{
			name:    "Should throw error for wrong command",
			args:    []string{"wrong"},
			want:    fmt.Sprintf("%s\n%s\n", testdata.WrongCommandMsg("wrong", backupCmd.Name()), testdata.BackupAccountsUsage),
			wantErr: true,
		},
		{
			name:    "Should throw error for no command arguments",
			args:    []string{},
			want:    fmt.Sprintf("%s\n%s\n", testdata.Atleast1ArgumentRequiredMsg, testdata.BackupAccountsUsage),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			rootCmd.SetOut(b)
			rootCmd.SetErr(b)
			rootCmd.SetArgs(append([]string{backupCmd.Name()}, tt.args...))

			if _, err := rootCmd.ExecuteC(); (err != nil) != tt.wantErr {
				t.Errorf("%sCmd.Execute() error = %v, wantErr %v", backupCmd.Name(), err, tt.wantErr)
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
