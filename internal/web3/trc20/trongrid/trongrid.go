package trongrid

import (
	"encoding/json"
	"fmt"
	"github.com/hipeday/upay/internal/util"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	domain string
}

func NewTronGridClient(domain string) *Client {
	return &Client{domain: domain}
}

// GetContractTransactions check from https://developers.tron.network/reference/get-transaction-info-by-account-address
func (t Client) GetContractTransactions(payload interface{}) (interface{}, error) {
	var (
		err  error
		path = "/v1/accounts/%s/transactions/trc20"
	)

	if payload == nil {
		return nil, fmt.Errorf("payload is nil")
	}
	transactionsPayload, ok := payload.(GetContractTransactionsPayload)
	if !ok {
		return nil, fmt.Errorf("payload is not GetContractTransactionsPayload")
	}

	u, err := url.Parse(fmt.Sprintf(t.domain+path, *transactionsPayload.Address))

	if err != nil {
		return nil, err
	}

	query := u.Query()

	params := util.BuildQueryParams(transactionsPayload)

	for key := range params {
		query.Add(key, params.Get(key))
	}

	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code %d", resp.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("failed to close response body")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response Response[[]ContractTransaction]

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
