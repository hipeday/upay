package trongrid

type Response[T any] struct {
	Data    *T   `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
	Meta    *struct {
		At          int64  `json:"at,omitempty" json:"at,omitempty"`
		PageSize    int64  `json:"page_size,omitempty" json:"page_size,omitempty"`
		Fingerprint string `json:"fingerprint,omitempty" json:"fingerprint,omitempty"`
		Links       struct {
			Next string `json:"next,omitempty" json:"next,omitempty" json:"next,omitempty"`
		} `json:"links" json:"links"`
	} `json:"meta"`
}

type ContractTransaction struct {
	TransactionId string `json:"transaction_id"`
	TokenInfo     struct {
		Symbol   string `json:"symbol"`
		Address  string `json:"address"`
		Decimals int    `json:"decimals"`
		Name     string `json:"name"`
	} `json:"token_info"`
	BlockTimestamp int64  `json:"block_timestamp"`
	From           string `json:"from"`
	To             string `json:"to"`
	Type           string `json:"type"`
	Value          string `json:"value"`
}
