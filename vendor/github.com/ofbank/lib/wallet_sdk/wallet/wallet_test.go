package wallet

import (
	"testing"
	"fmt"
	"math/big"
)

func TestGeneratePrivatekey(t *testing.T) {
	pri,err:=GeneratePrivatekey()
	if err!=nil{
		fmt.Println(err)
	}
fmt.Println(pri.D.String())
}


func TestSign(t *testing.T) {
	pri, err := GeneratePrivatekey()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pri.D.String())
	address := FromPrivateToAddress(pri, []int{1,2,3,4,5})
	fmt.Println("address: ",address,"private: ",pri.D.String())
	result,err:=Sign(address,
		address,
		new(big.Int).SetInt64(123),
		new(big.Int).SetInt64(123),
		new(big.Int).SetInt64(5000000000),
		3,
		[]byte{},
		new(big.Int).SetInt64(123),
		pri.D.String())
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(result)
}