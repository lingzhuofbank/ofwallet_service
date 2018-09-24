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