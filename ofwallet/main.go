package main

import (
	"flag"
	"github.com/ofwallet_service/utils"
	"github.com/ofwallet_service/api"
	"github.com/ofwallet_service/context"
	"github.com/ofwallet_service/utils/http"
)

var SERVER_PORT string

type rpc_service struct {
	name string
	method string
	description string
}

func main(){
	flag.StringVar(&SERVER_PORT, "port", "8085", "Http server port")
	flag.IntVar(&context.Acc,"acc",6,"token精度")
	utils.CreateLogger()
	flag.Parse()
	rpc :=http.GetRPCClient()
	for _,value:=range Services(){
		rpc.RegisterServices(value.name,value.method)
	}
	rpc.AddResquestAddress("http://47.92.99.227:8888")
	rpc.AddResquestAddress("http://47.92.173.173:8888")
	rpc.AddResquestAddress("http://47.92.105.199:8888")
	rpc.AddResquestAddress("http://47.92.99.227:8888")

	httpser := api.NewHttpServer(SERVER_PORT)
	httpser.StartServer()
}

func Services() []rpc_service {
	return []rpc_service{
		{
			name:        "sendTransaction",
			method:      "eth_sendRawTransaction",
			description: "发送交易",
		},
		{
			name:        "height",
			method:      "ofbank_lastBN",
			description: "获取区块高度",
		},
		{
			name:        "nonce",
			method:      "eth_getTransactionCount",
			description: "获取交易地址的Nonce值",
		},
		{
			name:        "balance",
			method:      "ofbank_show",
			description: "查询地址余额",
		},
		{
			name:"contract",
			method:"eth_call",
			description:"智能合约调用",
		},
	}
}