package trongrid

import (
	"log"
	"strconv"
	"testing"
)

func TestClient_GetTransactions(t1 *testing.T) {
	client := Client{domain: "https://nile.trongrid.io"}
	accountAddress := "TYYi4mhbkUwcezbnXP5UyKUbcS15Kx1qt8"
	contractAddress := "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj"
	onlyTo := true
	onlyConfirmed := true
	payload := GetContractTransactionsPayload{
		Address:         &accountAddress,
		OnlyConfirmed:   &onlyConfirmed,
		OnlyTo:          &onlyTo,
		OnlyFrom:        nil,
		Limit:           nil,
		Fingerprint:     nil,
		OrderBy:         nil,
		MinTimestamp:    nil,
		MaxTimestamp:    nil,
		ContractAddress: &contractAddress,
	}
	transactions, err := client.GetContractTransactions(payload)
	if err != nil {
		t1.Fatal(err)
	}

	response := transactions.(Response[[]ContractTransaction])
	log.Printf("transactions: %v", transactions)
	log.Printf(strconv.FormatBool(response.Success))
}
