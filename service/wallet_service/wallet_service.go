package wallet_service

import (
	"github.com/ofbank/lib/wallet_sdk/wallet"
	"errors"
	"github.com/ofwallet_service/utils"
	"github.com/ofwallet_service/params"
	"math/big"
	"github.com/ofwallet_service/utils/http"
	"strconv"
	"encoding/hex"
	"github.com/ofwallet_service/context"
)

var (
	walletError          = errors.New("1000:创建钱包失败")
	HegithError          = errors.New("1001:获取高度失败")
	NonceError           = errors.New("1002:获取钱包Nonce失败")
	SignError            = errors.New("1003:交易签名失败")
	BalanceError         = errors.New("1004:地址余额不足")
	SendTransactionError = errors.New("1005:发送交易失败")
	BalanceErr           = errors.New("1006:查询余额失败")
	DataErr              = errors.New("1007:智能合约解析失败")
)

func CreateWallet() (string, string, error) {
	key, err := wallet.GeneratePrivatekey()
	if err != nil {
		utils.ErrorLogger(err.Error())
	}
	address := wallet.FromPrivateToAddress(key, params.GetDouLiaoICCode())
	return key.D.String(), address, nil
}

func SendTransaction(from, to string, value, gasPrice, gasLimit *big.Int, private string, data string) (string, error) {
	height, err := GetHeight()
	if err != nil {
		return "", HegithError
	}
	nonce, err := getNonce(from)
	if err != nil {
		return "", NonceError
	}
	dataByte := []byte{}
	if data != "" &&len(data)>2{
		dataByte, err = hex.DecodeString(data[2:])
		if err != nil {
			utils.ErrorLogger("智能合约解析失败: "+err.Error())
			return "", DataErr
		}
	}else if data=="0x"{
		dataByte = []byte{}
	}

	siginedTransaction, err := wallet.Sign(from, to, value, gasLimit, gasPrice, nonce, dataByte, new(big.Int).SetInt64(height), private)
	if err != nil {
		utils.ErrorLogger("SginTransaction failed: " + err.Error())
		return "", SignError
	}
	txhash, err := sendTransaction(siginedTransaction)
	if err != nil {
		if err == BalanceError {
			return "", err
		} else {
			return "", SendTransactionError
		}
	}

	return txhash, nil
}

func CheckBalance(address string, contranctAddress string) (string,string,error) {
	balance, err := getBalance(address)
	if err != nil {
		utils.ErrorLogger("Failed to get balance: "+err.Error())
		return "","",BalanceErr
	}
	if contranctAddress==""{
		return balance,"0",nil
	}
	tokenBalance,err:=getTokenBalance(address,contranctAddress,context.Acc)
	if err!=nil{
		utils.ErrorLogger("Failed to get Token balance: "+err.Error())
		return "","",err
	}
	return balance,tokenBalance,nil
}

func getBalance(from string) (string, error) {

	var balance string
	err := http.GetRPCClient().Call("balance", &balance, from, "latest")
	if err != nil {
		return "", err
	}
	return balance, nil
}

func getNonce(from string) (uint64, error) {
	var nonceString string
	err := http.GetRPCClient().Call("nonce", &nonceString, from, "latest")
	if err != nil {
		return 0, err
	}
	nonce, err := strconv.ParseInt(nonceString, 0, 64)

	return uint64(nonce), nil
}



func sendTransaction(transaction string) (string, error) {
	var txhash string
	err := http.GetRPCClient().Call("sendTransaction", &txhash, transaction)
	if err != nil {
		return "", err
	}
	return txhash, nil
}

func getTokenBalance(from, contractAddress string, acc int) (string, error) {
	paramMap := make(map[string]string)
	paramMap["to"] = contractAddress
	paramMap["data"] = "0x70a08231" + "00000000000000" + from[2:]
	var contractResponse string
	err := http.GetRPCClient().Call("contract", &contractResponse, paramMap, "latest")
	if err != nil {
		return "", err
	}

	accPow := pow(10, acc)
	value, err := strconv.ParseInt(contractResponse, 0, 64)
	if err != nil {
		return "", err
	}
	valueResult := float64(value) / float64(accPow)
	floatString := strconv.FormatFloat(valueResult, 'f', acc, 64)
	return floatString, nil

}

func pow(x, n int) int {
	ret := 1
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
