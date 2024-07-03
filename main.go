package main

import (
	"fmt"
	"github.com/hipeday/upay/cmd/server"
	"io"
	"net/http"
	"sync"
)

const sandboxDomain = "https://nile.trongrid.io"

const getAccountInfoPath = "/v1/accounts/%v"
const getContractTransactionPath = "/v1/accounts/%v/transactions/trc20"

var once sync.Once

func main() {
	// setup server
	server.Run()
}

func getAccountInfo(address string) {

	address = "TYYi4mhbkUwcezbnXP5UyKUbcS15Kx1qt8"
	tokenAddress := "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj"
	getAccountInfo(address)
	getContractTransaction(tokenAddress, address)
	resp, err := http.Get(sandboxDomain + fmt.Sprintf(getAccountInfoPath, address))
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func getContractTransaction(contractAddress string, accountAddress string) {
	resp, err := http.Get(sandboxDomain + fmt.Sprintf(getContractTransactionPath, accountAddress) + "?contract_address=" + contractAddress)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
