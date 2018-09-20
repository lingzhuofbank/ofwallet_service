package wallet_service

import (
	"github.com/ofbank/lib/wallet_sdk/wallet"
	"errors"
	"github.com/ofwallet_service/utils"
	"github.com/ofwallet_service/params"
)


var (
	walletError=errors.New("1000:创建钱包失败")
)



func CreateWallet()(string,string,error){
	key,err:=wallet.GeneratePrivatekey()
	if err!=nil{
		utils.ErrorLogger(err.Error())
	}
    address:=wallet.FromPrivateToAddress(key,params.GetDouLiaoICCode())
    return key.D.String(),address,nil
}


func SendTransaction(from,to,value,private string) error{

}