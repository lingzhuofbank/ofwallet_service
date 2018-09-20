package main

import (
	"flag"
	"github.com/ofwallet_service/utils"
	"github.com/ofwallet_service/api"
)

var SERVER_PORT string

func main(){
	flag.StringVar(&SERVER_PORT, "port", "8085", "http server port")
	utils.CreateLogger()
	flag.Parse()
	httpser := api.NewHttpServer(SERVER_PORT)
	httpser.StartServer()
}
