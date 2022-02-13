package testdata

import (
	"fmt"
	"strings"

	"github.com/open-blockchain-explorer/tnbassist/model"
)

const (
	AccountNumber = "5decde0f7710e8baf449a0a1360babc1edee37e9e2020e81242882c8e51f3abc"
	SigningKey    = "0332c0639bd86982243a2dec82e2044dd587148fec1daf0c0cd5f22691c8916b"
	HelpText      = `TNB Assist is a CLI (Command Line Interface) tool for thenewboston blockchain to perform various mundane tasks like taking daily accounts backup, computing statistics, etc easier.

Checkout our flagship blockchain explorer at https://tnbexplorer.com

Usage:
  tnbassist [flags]
  tnbassist [command]

Available Commands:
  account     
  backup      Backup data of the given <entity>
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  stats       Stats command
  version     Current version of TNB Assist CLI.

Flags:
      --docs      Generate markdown docs for all commands
  -h, --help      help for tnbassist
  -v, --version   Print version information

Use "tnbassist [command] --help" for more information about a command.
`
	AccountVerifyUsage = `Usage:
  tnbassist account [flags]

Flags:
  -a, --acc string   Account number hex
  -h, --help         help for account
  -s, --sk string    Signing key hex
`

	BackupAccountsUsage = `Usage:
  tnbassist backup <entity> [flags]
  tnbassist backup [command]

Available Commands:
  verify      Verify data of the given <entity>.

Flags:
  -h, --help              help for backup
  -H, --host string       Base address/IP of primary validator server (required)
  -o, --outdir string     Output directory path where backup will be saved (default "./assets/account_backups")
  -p, --port uint16       Port of primary validator server (default 80)
      --save-as-json      Save to account backup in a JSON file
  -s, --scheme string     URL Scheme/Protocol of primary validator server (default "http")
  -t, --strftime string   String format time pattern used as filename for backup (default "%Y-%m-%d-%H_%M_%S")

Use "tnbassist backup [command] --help" for more information about a command.
`

	StatsOverviewUsage = `Usage:
  tnbassist stats overview [flags]

Flags:
      --append-to-csv       Path to account backup JSON file (make sure atleast 5, 10, 25 and 50 quantiles are provided)
  -b, --blacklist strings   Path to account backup JSON file
  -f, --file string         Path to account backup JSON file
  -h, --help                help for overview
  -H, --host string         Base address/IP of primary validator server (required)
      --json                Output should be in JSON format
  -o, --outdir string       Output directory path where backup will be saved (default "./assets")
  -p, --port uint16         Port of primary validator server (default 80)
      --post-to-server      Path to account backup JSON file (make sure TNBExplorer API token is provided)
  -q, --quantiles uints     Path to account backup JSON file (default [5,10,25,50])
  -s, --scheme string       URL Scheme/Protocol of primary validator server (default "http")
  -t, --token string        Path to account backup JSON file
`

	Atleast1ArgumentRequiredMsg = "Error: requires at least 1 arg(s), only received 0"
)

var FakeAccounts = model.Accounts{
	AccountNumber: model.AccountInfo{
		Balance:     0,
		BalanceLock: AccountNumber,
	},
}
var FakeAccountsJSON = fmt.Sprintf(`{
  "%s": {
    "balance": 0,
    "balance_lock": "%s"
  }
}
`, AccountNumber, AccountNumber)

func WrongCommandMsg(argument string, command string) string {
	return fmt.Sprintf("Error: invalid argument \"%s\" for \"tnbassist %s\"", argument, command)
}

func RequiredFlagMsg(flags ...string) string {
	return fmt.Sprintf(`required flag(s) "%s" not set`, strings.Join(flags, `", "`))
}
