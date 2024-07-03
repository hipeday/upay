package util

import (
	"github.com/hipeday/upay/internal/web3/trc20/trongrid"
	"log"
	"testing"
)

func TestBuildQueryParams(t *testing.T) {
	address := "address"
	onlyTo := true
	onlyConfirmed := true
	payload := trongrid.GetContractTransactionsPayload{
		Address:       &address,
		OnlyConfirmed: &onlyConfirmed,
		OnlyTo:        &onlyTo,
		OnlyFrom:      nil,
		Limit:         nil,
		Fingerprint:   nil,
		OrderBy:       nil,
		MinTimestamp:  nil,
		MaxTimestamp:  nil,
	}
	params := BuildQueryParams(payload)
	for s := range params {
		log.Println(s, params.Get(s))
	}
}
