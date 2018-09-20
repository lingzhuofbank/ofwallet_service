package wallet

import (
	"testing"
	"fmt"
)

func TestGeneratePrivatekey(t *testing.T) {
	pri,err:=GeneratePrivatekey()
	if err!=nil{
		fmt.Println(err)
	}
fmt.Println(pri.D.String())
}