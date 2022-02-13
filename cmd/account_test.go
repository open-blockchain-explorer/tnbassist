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

func Test_AccountCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
		regex   bool
	}{
		// Create
		{
			name:    "Should create account",
			args:    []string{"create"},
			want:    "Signing Key:\t\\s\\w+\nAccount No.:\t\\s\\w+\n",
			wantErr: false,
			regex:   true,
		},
		// Verify
		{
			name:    "Should return Valid if signing key actually belongs to the account",
			args:    []string{"verify", "-a", testdata.AccountNumber, "-s", testdata.SigningKey},
			want:    "Valid\n",
			wantErr: false,
		},
		{
			name:    "Should return Inalid if signing key actually belongs to the account",
			args:    []string{"verify", "-a", testdata.AccountNumber, "-s", testdata.AccountNumber},
			want:    "Invalid\n",
			wantErr: false,
		},
		{
			name:    "Should return error if signing key and account number is not provided",
			args:    []string{"verify"},
			want:    fmt.Sprintf("%s\n%s", testdata.AccountVerifyUsage, testdata.RequiredFlagMsg("sk", "acc")),
			wantErr: false,
		},
		{
			name:    "Should return error if signing key is not provided to verify account",
			args:    []string{"verify", "-a", testdata.AccountNumber},
			want:    fmt.Sprintf("%s\n%s", testdata.AccountVerifyUsage, testdata.RequiredFlagMsg("sk")),
			wantErr: false,
		},
		{
			name:    "Should return error if account number is not provided to verify account",
			args:    []string{"verify", "-s", testdata.SigningKey},
			want:    fmt.Sprintf("%s\n%s", testdata.AccountVerifyUsage, testdata.RequiredFlagMsg("acc")),
			wantErr: false,
		},
		// Restore
		{
			name:    "Should restore account for the given signing key to verify account",
			args:    []string{"restore", "-s", testdata.SigningKey},
			want:    fmt.Sprintf("Signing Key:\t %s\nAccount No.:\t %s\n", testdata.SigningKey, testdata.AccountNumber),
			wantErr: false,
		},
		{
			name:    "Should return error if signing key is not provided for restore",
			args:    []string{"restore"},
			want:    fmt.Sprintf("%s\n%s", testdata.AccountVerifyUsage, testdata.RequiredFlagMsg("sk")),
			wantErr: false,
		},
		// Wrong command
		{
			name:    "Should throw error for wrong command",
			args:    []string{"wrong"},
			want:    fmt.Sprintf("%s\n%s\n", testdata.WrongCommandMsg("wrong", accountCmd.Name()), testdata.AccountVerifyUsage),
			wantErr: true,
		},
		{
			name:    "Should throw error for no command arguments",
			args:    []string{},
			want:    fmt.Sprintf("%s\n%s\n", testdata.Atleast1ArgumentRequiredMsg, testdata.AccountVerifyUsage),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			rootCmd.SetOut(b)
			rootCmd.SetErr(b)
			rootCmd.SetArgs(append([]string{accountCmd.Name()}, tt.args...))

			if _, err := rootCmd.ExecuteC(); (err != nil) != tt.wantErr {
				t.Errorf("%sCmd.Execute() error = %v, wantErr %v", accountCmd.Name(), err, tt.wantErr)
			}
			accountCmd.Flags().VisitAll(func(flag *pflag.Flag) {
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
