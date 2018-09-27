package http

import (
	"sync"
	"sync/atomic"
	"encoding/json"
	"errors"
	"reflect"
	"github.com/ofwallet_service/utils"

)

var rpcClient *RPCClient

type RPCRequest struct {
	Id     uint32         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type RPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Result  interface{} `json:"result"`
	Error   *RPCError   `json:"error"`
}

type RPCError struct {
	Code    int
	Message string
}

type RPCClient struct {
	idCounter      uint32
	Service        map[string]string
	RequestAddress []string
	CurrentIndex   int
	indexLock      *sync.Mutex
}

func newRPCClient() *RPCClient {
	return &RPCClient{
		Service:        make(map[string]string),
		RequestAddress: make([]string, 0),
		CurrentIndex:   0,
		indexLock:      new(sync.Mutex),
	}
}

func (rc *RPCClient) GetAddress() string {

	defer func() {
		rc.indexLock.Lock()
		if rc.CurrentIndex == len(rc.RequestAddress)-1 {
			rc.CurrentIndex = 0
		} else {
			rc.CurrentIndex++
		}
		rc.indexLock.Unlock()
	}()

	return rc.RequestAddress[rc.CurrentIndex]
}

func (rc *RPCClient) AddResquestAddress(address string) {
	rc.RequestAddress = append(rc.RequestAddress, address)
}

func (rc *RPCClient) RegisterServices(name, service string) {
	rc.Service[name] = service
}

func GetRPCClient() *RPCClient {
	if rpcClient == nil {
		rpcClient = newRPCClient()
	}
	return rpcClient
}

func (rc *RPCClient) Call(methodName string, response interface{}, params ...interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			utils.ErrorLogger(err)
		}
	}()
	rpcRequest := RPCRequest{
		Id:     rc.nextID(),
		Method: rc.Service[methodName],
		Params: params,
	}
	address := rc.GetAddress()
	responseByte, err := SendRequest(address, POST, rpcRequest)
	if err != nil {
		return err
	}
	var rpcReponse RPCResponse
	err = json.Unmarshal(responseByte, &rpcReponse)
	if err != nil {
		utils.ErrorLogger("RPC解析json失败：" + err.Error())
		return err
	}

	if rpcReponse.Error != nil {
		return errors.New(rpcReponse.Error.Message)
	}
	if rpcReponse.Result==nil{
		response=nil
		return nil
	}
	err = setValue(response, rpcReponse.Result)
	if err != nil {
		return err
	}

	return nil
}

func setValue(v interface{}, value interface{}) error {
	switch v.(type) {
	case *string:
		reflect.ValueOf(v).Elem().SetString(value.(string))
		return nil
	case *int, *int64:
		reflect.ValueOf(v).Elem().SetInt(value.(int64))
		return nil
	case *float64:
		reflect.ValueOf(v).Elem().SetFloat(value.(float64))
		return nil
	case *bool:
		reflect.ValueOf(v).Elem().SetBool(value.(bool))
		return nil
	case []byte:
		byteResult, err := json.Marshal(value)
		if err != nil {
			return err
		}
		reflect.ValueOf(v).Elem().SetBytes(byteResult)
	case interface{}:
		result, err := json.Marshal(value)
		if err != nil {
			return err
		}
		err = json.Unmarshal(result, v)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (rc *RPCClient) nextID() uint32 {
	id := atomic.AddUint32(&rc.idCounter, 1)
	return id
}