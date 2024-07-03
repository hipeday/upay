package trongrid

type GetContractTransactionsPayload struct {
	Address         *string `query:"address,skip"`
	OnlyConfirmed   *bool   `query:"only_confirmed"`
	OnlyUnconfirmed *bool   `query:"only_unconfirmed"`
	OnlyTo          *bool   `query:"only_to"`
	OnlyFrom        *bool   `query:"only_from"`
	Limit           *int32  `query:"limit"`
	Fingerprint     *string `query:"fingerprint"`
	OrderBy         *string `query:"order_by"`
	MinTimestamp    *string `query:"min_timestamp"`
	MaxTimestamp    *string `query:"max_timestamp"`
	ContractAddress *string `query:"contract_address"`
}
