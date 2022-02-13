package compute

import (
	"log"
	"os"

	"github.com/bcicen/jstream"
	"github.com/mitchellh/mapstructure"
	"github.com/open-blockchain-explorer/tnbassist/model"
)

func streamAccounts(file *os.File, ch chan<- model.Account) {
	decoder := jstream.NewDecoder(file, 1).EmitKV() // extract JSON values at a depth level of 1
	for mv := range decoder.Stream() {
		switch t := mv.Value.(type) {
		case jstream.KV:
			var account model.Account
			if err := mapstructure.Decode(t.Value, &account); err != nil {
				log.Fatal(err)
			}
			account.AccountNumber = t.Key
			ch <- account
		default:
			log.Fatalf("Failed to parse (%+v) as JSON properly", t)
		}
	}
	close(ch)
}
