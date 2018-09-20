package api

import (
	"net/http"
	"github.com/ofwallet_service/filter"
	"github.com/ofwallet_service/api/wallet"
	"os"
	"github.com/ofwallet_service/utils"
)

type HttpServer struct {
	hs   *http.Server
	Done chan struct{}
}

func NewHttpServer(port string) *HttpServer {

	http.HandleFunc("/ofbank/address", filter.ServiceFilterMapper(wallet.CreateAddress))

	server := new(http.Server)
	server.Addr = ":" + port
	server.Handler = nil

	return &HttpServer{
		hs:   server,
		Done: make(chan struct{}),
	}
}


func (hs *HttpServer) StartServer() {
	utils.NoticeLooger("Start API Service: "+hs.hs.Addr)
	err := hs.hs.ListenAndServe()

	if err != nil {
		utils.ErrorLogger("Failed to set up services: "+err.Error() )
		os.Exit(1)
	}
}
