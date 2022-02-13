//go:build !mock
// +build !mock

package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/open-blockchain-explorer/tnbassist/model"
)

// TNBHTTPClient implements TNBHTTPClientInterface
// It provides methods to access thenewboston services via RESTFul APIs
type TNBHTTPClient struct {
	BaseAddress  HTTPConfig
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// NewTNBHTTPClient instantiates new TNBHTTPClientInterface compatible client
func NewTNBHTTPClient(httpConfig HTTPConfig) TNBHTTPClientInterface {
	return &TNBHTTPClient{
		BaseAddress:  httpConfig,
		readTimeout:  5 * time.Second,
		writeTimeout: 10 * time.Second,
	}
}

// FetchAllAccounts fetches all accounts from thenewboston
func (t *TNBHTTPClient) FetchAllAccounts() (*model.Accounts, error) {
	response := model.Accounts{}

	url := fmt.Sprintf("%s/accounts?limit=%d&offset=%d", t.BaseAddress, 0, 0)
	method := "GET"

	client := &http.Client{
		Timeout: t.readTimeout,
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var accountsList []model.Account
	err = json.Unmarshal(body, &accountsList)
	if err != nil {
		return nil, err
	}

	for _, account := range accountsList {
		response[account.AccountNumber] = model.AccountInfo{
			Balance:     account.Balance,
			BalanceLock: account.BalanceLock,
		}
	}

	return &response, nil
}

// FetchAccountsWithLimitAndOffset fetches upto 100 accounts from thenewboston
func (t *TNBHTTPClient) FetchAccountsWithLimitAndOffset(limit uint8, offset uint) (*model.PaginatedAccounts, error) {
	var response model.PaginatedAccounts

	if limit > 100 {
		limit = 100
	}

	url := fmt.Sprintf("%s/accounts?limit=%d&offset=%d", t.BaseAddress, limit, offset)
	method := "GET"

	client := &http.Client{
		Timeout: t.readTimeout,
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
