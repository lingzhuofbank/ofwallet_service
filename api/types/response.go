package types

type ResponseModel struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

type CreateAddressResponse struct {
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
}

type BalacneResponse struct {
	OFBalance  string 	`json:"of_balance"`
	TokenBalance string  `json:"token_balance"`
}

type TransactionResponse struct {
	BlockHash string  `json:"blockHash"`
	From      string  `json:"from"`
	Gas       string  `json:"gas"`
	GasPrice  string  `json:"gasPrice"`
	Hash      string  `json:"hash"`
	Input     string  `json:"input"`
	Nonce     string     `json:"nonce"`
	Timestamp string  `json:"timestamp"`
	TxFee     string  `json:"txFee"`
	Value     string `json:"value"`
}