package model

import (
	"fmt"
)

// LegacyStats is a struct that represents the stats in Legacy structure
type LegacyStats struct {
	Date           string  `json:"date" csv:"Date"`
	Shift          uint64  `json:"shift" csv:"Shift"`
	Total          uint64  `json:"total" csv:"Total"`
	NAccounts      uint32  `json:"accounts" csv:"Accounts"`
	MaxBalance     uint64  `json:"max_balance" csv:"Max balance"`
	RichestAccount string  `json:"richest" csv:"Richest"`
	Top5Ownership  float32 `json:"top_5_ownership" csv:"Top 5% ownership"`
	Top5NAccounts  uint32  `json:"top_50_wealth" csv:"Top 5% accounts"`
	Top5Wealth     uint64  `json:"top_50_accounts" csv:"Top 5% wealth"`
	Top10Ownership float32 `json:"top_50_ownership" csv:"Top 10% ownership"`
	Top10NAccounts uint32  `json:"top_25_wealth" csv:"Top 10% accounts"`
	Top10Wealth    uint64  `json:"top_25_accounts" csv:"Top 10% wealth"`
	Top25Ownership float32 `json:"top_25_ownership" csv:"Top 25% ownership"`
	Top25NAccounts uint32  `json:"top_10_wealth" csv:"Top 25% accounts"`
	Top25Wealth    uint64  `json:"top_10_accounts" csv:"Top 25% wealth"`
	Top50Ownership float32 `json:"top_10_ownership" csv:"Top 50% ownership"`
	Top50NAccounts uint32  `json:"top_5_wealth" csv:"Top 50% accounts"`
	Top50Wealth    uint64  `json:"top_5_accounts" csv:"Top 50% wealth"`
}

// Stats represents the stats of the blockchain network
type Stats struct {
	Shift          uint64              `json:"shift"`
	Total          uint64              `json:"total"`
	NAccounts      uint32              `json:"accounts"`
	MaxBalance     uint64              `json:"max_balance"`
	RichestAccount string              `json:"richest"`
	Quanitiles     map[uint8]Quanitile `json:"quantiles"`
	RichList       []Rich              `json:"richlist"`
}

// Rich represents individual account and its balance
type Rich struct {
	Account string `json:"account"`
	Balance uint64 `json:"balance"`
}

// Quanitile represents the distribution of wealth and ownership for specific percentile of the richlist
type Quanitile struct {
	Ownership float32 `json:"ownership"`
	NAccounts uint32  `json:"accounts"`
	Wealth    uint64  `json:"wealth"`
}

// AddQuantile adds a new quanitile to the stats
func (s *Stats) AddQuantile(pct uint8, quantile Quanitile) *Stats {
	if s.Quanitiles == nil {
		s.Quanitiles = make(map[uint8]Quanitile)
	}
	if pct > 100 {
		return s
	}
	s.Quanitiles[pct] = quantile
	return s
}

// ToLegacyStats converts the stats to the legacy stats
func (s *Stats) ToLegacyStats(date string) *LegacyStats {
	ls := LegacyStats{
		Date:           date,
		Shift:          s.Shift,
		Total:          s.Total,
		NAccounts:      s.NAccounts,
		MaxBalance:     s.MaxBalance,
		RichestAccount: s.RichestAccount,
	}

	if val, ok := s.Quanitiles[5]; ok {
		ls.Top5NAccounts = val.NAccounts
		ls.Top5Ownership = val.Ownership
		ls.Top5Wealth = val.Wealth
	}
	if val, ok := s.Quanitiles[10]; ok {
		ls.Top10NAccounts = val.NAccounts
		ls.Top10Ownership = val.Ownership
		ls.Top10Wealth = val.Wealth
	}
	if val, ok := s.Quanitiles[25]; ok {
		ls.Top25NAccounts = val.NAccounts
		ls.Top25Ownership = val.Ownership
		ls.Top25Wealth = val.Wealth
	}
	if val, ok := s.Quanitiles[50]; ok {
		ls.Top50NAccounts = val.NAccounts
		ls.Top50Ownership = val.Ownership
		ls.Top50Wealth = val.Wealth
	}
	return &ls
}

// CSVRow returns the stats as a slice of strings
func (l *LegacyStats) CSVRow() []string {
	return []string{
		l.Date,
		fmt.Sprintf("%d", l.Shift),
		fmt.Sprintf("%d", l.Total),
		fmt.Sprintf("%d", l.NAccounts),
		fmt.Sprintf("%d", l.MaxBalance),
		l.RichestAccount,
		fmt.Sprintf("%d", l.Top5Wealth),
		fmt.Sprintf("%.2f", l.Top5Ownership),
		fmt.Sprintf("%d", l.Top5NAccounts),
		fmt.Sprintf("%d", l.Top10Wealth),
		fmt.Sprintf("%.2f", l.Top10Ownership),
		fmt.Sprintf("%d", l.Top10NAccounts),
		fmt.Sprintf("%d", l.Top25Wealth),
		fmt.Sprintf("%.2f", l.Top25Ownership),
		fmt.Sprintf("%d", l.Top25NAccounts),
		fmt.Sprintf("%d", l.Top50Wealth),
		fmt.Sprintf("%.2f", l.Top50Ownership),
		fmt.Sprintf("%d", l.Top50NAccounts),
	}
}

// CSVHeaders returns the header for the stats
func (l *LegacyStats) CSVHeaders() []string {
	return []string{
		"Date",
		"Shift",
		"Total",
		"Accounts",
		"Max balance",
		"Richest",
		"Top 5% wealth",
		"Top 5% ownership",
		"Top 5% accounts",
		"Top 10% wealth",
		"Top 10% ownership",
		"Top 10% accounts",
		"Top 25% wealth",
		"Top 25% ownership",
		"Top 25% accounts",
		"Top 50% wealth",
		"Top 50% ownership",
		"Top 50% accounts",
	}
}
