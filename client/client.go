/*
Package client provides a client for the HTTP API.
*/
package client

import (
	"fmt"

	"github.com/open-blockchain-explorer/tnbassist/model"
)

// HTTPConfig represents HTTP URI elements.
type HTTPConfig struct {
	Protocol string
	Host     string
	Port     uint16
}

// String returns URI string representation of HTTPConfig
func (h HTTPConfig) String() string {
	return fmt.Sprintf("%s://%s:%d", h.Protocol, h.Host, h.Port)
}

// TNBExplorerHTTPClientInterface represents the interface for TNBExplorer RESTful API client
//
// Interfaces serve as parent objects from which other objects can inherit.
// It guarantees that a client provides basic functionality.
// Interface can be easily mocked making it easier to test the client.
type TNBExplorerHTTPClientInterface interface {
	PostStats(stats *model.LegacyStats) (int, []byte, error)
}

// TNBHTTPClientInterface represents the interface for thenewboston RESTful API client
//
// Interfaces serve as parent objects from which other objects can inherit.
// It guarantees that a client provides basic functionality.
// Interface can be easily mocked making it easier to test the client.
type TNBHTTPClientInterface interface {
	FetchAllAccounts() (*model.Accounts, error)
	FetchAccountsWithLimitAndOffset(limit uint8, offset uint) (*model.PaginatedAccounts, error)
}

// DiscordWebhookClientInterface represents the interface for Discord RESTful API client
//
// Interfaces serve as parent objects from which other objects can inherit.
// It guarantees that a client provides basic functionality.
// Interface can be easily mocked making it easier to test the client.
type DiscordWebhookClientInterface interface {
	PostStatsToDiscordFetchAllAccounts() (int, error)
}

// GitHubGraphQLClientInterface represents the interface for GitHub GraphQL client
//
// Interfaces serve as parent objects from which other objects can inherit.
// It guarantees that a client provides basic functionality.
// Interface can be easily mocked making it easier to test the client.
type GitHubGraphQLClientInterface interface {
	FetchNIssues(organization string, repository string, filters Filters) (*Issues, error)
	CheckRateLimit() (*RateLimit, error)
	WhoAmI() (*Viewer, error)
}
