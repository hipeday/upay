package web3

type Client interface {
	// GetContractTransactions 通过账户地址获取合约交易信息
	GetContractTransactions(payload interface{}) (interface{}, error)
}
