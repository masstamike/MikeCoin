package main

import (
	"net/http"
	"blockchain/api"
)

func main() {

	var mc api.MikeCoin
	mc.Initialize()

	mux := http.DefaultServeMux
	mux.HandleFunc("/balance", mc.Balance)
	mux.HandleFunc("/users", mc.Users)
	mux.HandleFunc("/transfers", mc.Transfers)
	mux.HandleFunc("/status", mc.GetStatus)
	http.ListenAndServe(":9090", mux) // set listen port
}
