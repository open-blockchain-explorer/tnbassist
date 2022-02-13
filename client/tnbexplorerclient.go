package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/open-blockchain-explorer/tnbassist/model"
)

// TNBExplorerHTTPClient implements TNBExplorerHTTPClientInterface
// It provides methods to access TNBExplorer services via RESTFul APIs
type TNBExplorerHTTPClient struct {
	BaseAddress  HTTPConfig
	token        string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// NewTNBExplorerHTTPClient instantiates new TNBExplorerHTTPClientInterface compatible client
func NewTNBExplorerHTTPClient(httpConfig HTTPConfig, token string) TNBExplorerHTTPClientInterface {
	return &TNBExplorerHTTPClient{
		BaseAddress:  httpConfig,
		token:        token,
		readTimeout:  5 * time.Second,
		writeTimeout: 10 * time.Second,
	}
}

// PostStats sends stats to TNBExplorer
func (t *TNBExplorerHTTPClient) PostStats(stats *model.LegacyStats) (int, []byte, error) {
	url := fmt.Sprintf("%s/stats/api/", t.BaseAddress)
	method := "POST"

	payload, err := json.Marshal(stats)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	client := &http.Client{
		Timeout: t.writeTimeout,
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", t.token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return res.StatusCode, body, nil
}
