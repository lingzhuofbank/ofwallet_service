package wallet_service

import (
	"strconv"
	"github.com/ofwallet_service/utils/http"
)


func GetHeight() (int64, error) {
	var heightString string
	err := http.GetRPCClient().Call("height", &heightString, nil)
	if err != nil {
		return -1, err
	}
	height, err := strconv.ParseInt(heightString, 0, 64)
	if err != nil {
		return -1, err
	}
	return height, nil
}
