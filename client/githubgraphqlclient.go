package client

import (
	"context"
	"time"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

// GitHubGraphqlURL is github's GraphQL API server.
const GitHubGraphqlURL = "https://api.github.com/graphql"

// MaxPagination is github's max pagination limit.
const MaxPagination = 100

// Open is an issue that is still open.
const Open IssueState = "OPEN"

// Closed is an issue that has been closed.
const Closed IssueState = "CLOSED"

// IssueState represents the possible states of an issue.
type IssueState string

// DateTime represents the time.
type DateTime struct {
	time.Time `json:"time,omitempty"`
}

// IssueFilters represents more granular ways to filter lists of issues.
type IssueFilters struct {
	Since *DateTime `json:"since,omitempty"`
}

// Filters represents ways to filter lists of issues.
type Filters struct {
	Limit      int8         `json:"limit,omitempty"`
	Offset     *string      `json:"offset,omitempty"`
	IssueState IssueState   `json:"issueState,omitempty"`
	Labels     []string     `json:"labels,omitempty"`
	FilterBy   IssueFilters `json:"filterBy,omitempty"`
}

// PageInfo represents information about pagination.
type PageInfo struct {
	StartCursor string `json:"startCursor,omitempty"`
	HasNextPage bool   `json:"hasNextPage,omitempty"`
	EndCursor   string `json:"endCursor,omitempty"`
}

// Issue represents a single issue from the repository.
type Issue struct {
	Title    string    `json:"title,omitempty"`
	Body     string    `json:"body,omitempty"`
	ClosedAt *DateTime `json:"closedAt,omitempty"`
	BodyURL  string    `json:"bodyUrl,omitempty"`
}

// Issues represents list of issues that have been created in the repository.
type Issues struct {
	PageInfo PageInfo `json:"pageInfo,omitempty"`
	Edges    []struct {
		Node Issue `json:"issue,omitempty"`
	} `json:"edges,omitempty"`
}

// Viewer represents the currently authenticated user.
type Viewer struct {
	Login     string    `json:"login,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt *DateTime `json:"createdAt,omitempty"`
}

// RateLimit represents the client's rate limit information.
type RateLimit struct {
	Limit     uint32    `json:"limit,omitempty"`
	Cost      uint32    `json:"cost,omitempty"`
	Remaining uint32    `json:"remaining,omitempty"`
	ResetAt   *DateTime `json:"resetAt,omitempty"`
}

// GitHubGraphQLClient implements GitHubGraphQLClientInterface
// It provides methods to access GitHub services via GraphQL APIs
type GitHubGraphQLClient struct {
	client graphql.Client
}

// NewGitHubGraphQLClient instantiates new GitHubGraphQLClientInterface compatible client
func NewGitHubGraphQLClient(token string) GitHubGraphQLClientInterface {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := graphql.NewClient(GitHubGraphqlURL, httpClient)

	return &GitHubGraphQLClient{
		client: *client,
	}
}

// FetchNIssues returns list of issues that have been created in the repository.
func (g GitHubGraphQLClient) FetchNIssues(organization string, repository string, filters Filters) (*Issues, error) {
	// avoid EXCESSIVE_PAGINATION error
	if filters.Limit > MaxPagination {
		filters.Limit = MaxPagination
	}
	var query struct {
		Repository struct {
			Issues Issues `graphql:"issues(first: $limit, after: $offset, states: $states, labels: $labels, filterBy: $filterBy)"`
		} `graphql:"repository(owner: $organization, name: $repository)"`
	}

	labels := []graphql.String{}
	for _, label := range filters.Labels {
		labels = append(labels, graphql.String(label))
	}

	variables := map[string]interface{}{
		"organization": graphql.String(organization),
		"repository":   graphql.String(repository),
		"limit":        graphql.Int(filters.Limit),
		"offset":       (*graphql.String)(filters.Offset),
		"states":       []IssueState{filters.IssueState},
		"filterBy":     filters.FilterBy,
		"labels":       labels,
	}
	err := g.client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	return &query.Repository.Issues, nil
}

// CheckRateLimit returns the client's rate limit information.
func (g GitHubGraphQLClient) CheckRateLimit() (*RateLimit, error) {
	var query struct {
		RateLimit RateLimit
	}
	err := g.client.Query(context.Background(), &query, nil)
	if err != nil {
		return nil, err
	}
	return &query.RateLimit, nil
}

// WhoAmI returns the currently authenticated user.
func (g GitHubGraphQLClient) WhoAmI() (*Viewer, error) {
	var query struct {
		Viewer Viewer
	}
	err := g.client.Query(context.Background(), &query, nil)
	if err != nil {
		return nil, err
	}
	return &query.Viewer, nil
}
