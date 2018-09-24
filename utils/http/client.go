package http

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"time"

	"github.com/ofwallet_service/utils"
	"strconv"
)


const (
	GET  = iota
	POST
)

func SendRequest(requestUrl string, method int, arg interface{}) ([]byte, error) {
	req := &fasthttp.Request{}
	switch argType := arg.(type) {
	case RPCRequest:
		bodyJson, _ := json.Marshal(argType)
		req.SetBody(bodyJson)
		req.Header.SetMethod("POST")
	case *fasthttp.Args:
		switch method {
		case POST:
			req.Header.SetMethod("POST")
			argType.WriteTo(req.BodyWriter())
		case GET:
			req.Header.SetMethod("GET")
			requestUrl = requestUrl + "?" + argType.String()
		}
	}

	req.SetRequestURI(requestUrl)
	resp := &fasthttp.Response{}
	err := fasthttp.DoTimeout(req, resp, time.Second*10)
	if err != nil {
		utils.ErrorLogger("Send Http Failed: "+err.Error())
		return nil, err
	}
	if resp.StatusCode() != 200 {
		utils.NoticeLooger("Send Http statusCode incorrect: "+strconv.Itoa(resp.StatusCode()))
	}

	return resp.Body(), nil
}
