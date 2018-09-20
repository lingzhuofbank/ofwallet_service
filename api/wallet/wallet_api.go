package wallet

import (
	"net/http"
	"github.com/ofwallet_service/api/types"

	"github.com/ofwallet_service/service/wallet_service"
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

//func SendTransaction(w http.ResponseWriter, r *http.Request) (*types.ResponseModel, error) {
//	//key := r.FormValue("private_key")
//	//value := r.FormValue("value")
//	//to := r.FormValue("to")
//	//from := r.FormValue("from")
//	//
//
//
//}
