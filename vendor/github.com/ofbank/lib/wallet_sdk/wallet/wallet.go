package wallet

import (
   "crypto/ecdsa"
   "github.com/ethereum/go-ethereum/crypto"

   "crypto/rand"
   "ofbc/common/hexutil"
)

func GeneratePrivatekey()(*ecdsa.PrivateKey,error){
   privateKey,err:=ecdsa.GenerateKey(crypto.S256(),rand.Reader)
   if err!=nil{
      return nil,err
   }
   return privateKey,nil
}


func FromPrivateToAddress(key *ecdsa.PrivateKey,iccode []int)string{
    address:=crypto.PubkeyToAddress(key.PublicKey,iccode)
    return hexutil.Encode(address.Bytes())
}