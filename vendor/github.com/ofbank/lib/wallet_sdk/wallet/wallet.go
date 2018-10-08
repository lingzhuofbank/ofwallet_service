package wallet

import (
   "crypto/ecdsa"
   "github.com/ethereum/go-ethereum/crypto"

   "crypto/rand"
   "github.com/ethereum/go-ethereum/core/types"
   "math/big"
   "github.com/ethereum/go-ethereum/common"
   "errors"
   "fmt"
   "github.com/ethereum/go-ethereum/rlp"
   "github.com/ethereum/go-ethereum/common/hexutil"
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

func Sign(from,to string, amount, gasLimit, gasPrice *big.Int, nonce uint64, data []byte, number *big.Int, prikey string) (string, error) {
   toByte, err := hexutil.Decode(to)
   if err != nil {
      return "", err
   }
   fromByte, err := hexutil.Decode(from)
   if err != nil {
      return "", err
   }

   abt := fromByte[:5]
   var cbt = []byte{0, 0, 0, 0, 156}
   cbt[0] = byte(abt[0])
   cbt[1] = byte(abt[1])
   cbt[3] = byte(abt[3])
   cbt[2] = byte(abt[2])
   cbt[4] = byte(abt[4])

   fmt.Println(cbt)

   privateKeyD, _ := new(big.Int).SetString(prikey, 10)
   prikeyKey, err := crypto.ToECDSA(privateKeyD.Bytes())
   if err != nil {
      return "", nil
   }

   recoverAddress:=FromPrivateToAddress(prikeyKey,[]int{int(cbt[0]),int(cbt[1]),int(cbt[2]),int(cbt[3]),int(cbt[4])})
   if recoverAddress==""||(recoverAddress!=from){
      return "",errors.New("钱包签名和转出地址不匹配")
   }


   transaction := types.NewTransaction(nonce, common.BytesToAddress(toByte), amount, gasLimit, gasPrice, nil, data, cbt)
   var sign types.Signer

   sign = types.HomesteadSigner{}

   transaction, err = types.SignTx(transaction, sign, prikeyKey)

   if err != nil {
      return "", nil
   }

   rlpBytes, err := rlp.EncodeToBytes(transaction)
   if err != nil {
      return "", err
   }
   return  hexutil.Encode(rlpBytes), nil
}
