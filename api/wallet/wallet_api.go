package wallet

import (
	"net/http"
	"github.com/ofwallet_service/api/types"

	"github.com/ofwallet_service/service/wallet_service"
	"math/big"
)

func CreateAddress(w http.ResponseWriter, r *http.Request) (*types.ResponseModel, error) {
	pri, addr, err := wallet_service.CreateWallet()
	if err != nil {
		return nil, err
	}
	response := &types.ResponseModel{
		Message: "创建地址成功",
		Code:    200,
		Data: &types.CreateAddressResponse{
			pri,
			addr,
		},
	}
	return response, nil
}

func SendTransaction(w http.ResponseWriter, r *http.Request) (*types.ResponseModel, error) {
	transactionRequest := &types.SendTransaction{}
	err := types.GetStructFromRequest(r, transactionRequest)
	if err != nil {
		return nil, err
	}
	amount, _ := new(big.Int).SetString(transactionRequest.Value, 10)
	GasPrice, _ := new(big.Int).SetString(transactionRequest.GasPrice, 10)
	txhash, err := wallet_service.SendTransaction(transactionRequest.From,
		transactionRequest.To,
		amount,
		GasPrice, new(big.Int).SetInt64(transactionRequest.GasLimit),
		transactionRequest.PrivateKey,
		transactionRequest.Data)
	if err != nil {
		return nil, err
	}

	reponse := &types.ResponseModel{
		Message: "发送交易成功",
		Code:    200,
		Data:    txhash,
	}
	return reponse, nil
}

func CheckBalce(w http.ResponseWriter, r *http.Request) (*types.ResponseModel, error) {
	address := r.FormValue("address")
	contranctAddress := r.FormValue("contract_address")
	balance, tokenBalance, err := wallet_service.CheckBalance(address, contranctAddress)
	if err != nil {
		return nil, err
	}
	reponse := &types.ResponseModel{
		Message: "查询余额成功",
		Code:    200,
		Data: &types.BalacneResponse{
			balance,
			tokenBalance,
		},
	}

	return reponse, nil
}
