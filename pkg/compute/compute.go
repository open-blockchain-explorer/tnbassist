package compute

import (
	"math"
	"os"

	"github.com/open-blockchain-explorer/tnbassist/model"
	"github.com/wangjia184/sortedset"
)

// Options for computing quantiles
type Options struct {
	Blacklist        map[string]struct{}
	Quantiles        []uint
	MaxRichListCount uint
}

// CoumputeQuantiles computes quantiles for a given account backup JSON file
func CoumputeQuantiles(file string, opts *Options) (*model.Stats, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ch := make(chan model.Account)
	go streamAccounts(f, ch)

	sortedAccounts := sortedset.New()

	var totalBalance uint64
	for v := range ch {
		if _, ok := opts.Blacklist[v.AccountNumber]; !ok {
			totalBalance += v.Balance
			// The rank is 1-based, that is to say, rank 1 is the node with minimum score
			sortedAccounts.AddOrUpdate(v.AccountNumber, sortedset.SCORE(v.Balance), nil)
		}
	}

	count := sortedAccounts.GetCount()

	stats := model.Stats{
		Total:          totalBalance,
		NAccounts:      uint32(count),
		MaxBalance:     uint64(sortedAccounts.PeekMax().Score()),
		RichestAccount: sortedAccounts.PeekMax().Key(),
		RichList:       make([]model.Rich, opts.MaxRichListCount),
	}

	maxPct := 1
	for _, pct := range opts.Quantiles {
		if pct <= 100 {
			stats.AddQuantile(uint8(pct), model.Quanitile{
				NAccounts: uint32(math.Floor(float64(pct) / 100 * float64(count))),
			})
			if int(pct) > maxPct {
				maxPct = int(pct)
			}
		}
	}
	minAccountScanNeeded := count / (100 / int(maxPct))
	if int(opts.MaxRichListCount) > minAccountScanNeeded {
		minAccountScanNeeded = int(opts.MaxRichListCount)
	}

	for idx, node := range sortedAccounts.GetByRankRange(-1, -minAccountScanNeeded, false) {
		if idx < int(opts.MaxRichListCount) {
			stats.RichList[idx] = model.Rich{
				Account: node.Key(),
				Balance: uint64(node.Score()),
			}
		}

		for pct, quantile := range stats.Quanitiles {
			// fmt.Println(idx, pct, quantile.NAccounts)
			if idx < int(quantile.NAccounts) {
				quantile.Wealth += uint64(node.Score())
				stats.Quanitiles[uint8(pct)] = quantile
			}
		}
	}

	for pct, quantile := range stats.Quanitiles {
		quantile.Ownership = float32(quantile.Wealth) / float32(totalBalance)
		stats.Quanitiles[uint8(pct)] = quantile
	}

	return &stats, nil
}
