package filter

import (
	"net/http"
	"github.com/ofwallet_service/api/types"
	"encoding/json"
	"github.com/ofwallet_service/utils"
	"strings"
	"strconv"
)

type serverHandler func(write http.ResponseWriter, request *http.Request) (*types.ResponseModel,error)

func ServiceFilterMapper(handler serverHandler) func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		utils.Logger.Notice("Request URL: ", request.URL.String())
		data,err:= handler(w, request)
		if err != nil {
			GenerateErrResponse(w, err)
		} else {
			GenerateSuccessResponse(w, data)
		}
	}
}

func GenerateErrResponse(write http.ResponseWriter, err error) {
	defer func() {
		if err:=recover();err!=nil{
			utils.ErrorLogger(err)
		}
	}()
	ErrMessage := err.Error()
	ErrArr := strings.Split(ErrMessage, ":")
	status, _ := strconv.Atoi(ErrArr[0])
	model := &types.ResponseModel{
		Code:    status,
		Message: ErrArr[1],
		Data:    nil,
	}
	jsonByte, _ := json.Marshal(model)
	utils.Logger.Error("Request Response ERROR:", string(jsonByte))
	write.WriteHeader(200)
	write.Write(jsonByte)
}

func GenerateSuccessResponse(write http.ResponseWriter, model *types.ResponseModel) {
	resp, _ := json.Marshal(model)
	utils.Logger.Notice("Request Response SUCCESS:", string(resp))
	write.WriteHeader(200)
	write.Write(resp)
}
