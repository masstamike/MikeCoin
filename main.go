package main

import (
	"net/http"
	"blockchain/api"
	"blockchain/api/gossip"
	"os"
	"strconv"
	"fmt"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Incorrect number of args, expects 2.")
		os.Exit(1)
	}
	myPort := os.Args[1]
	peerPort, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var mc api.MikeCoin
	mc.Initialize()
	gossip.Initialize(peerPort)


	mux := http.DefaultServeMux
	mux.HandleFunc("/balance", mc.Balance)
	mux.HandleFunc("/users", mc.Users)
	mux.HandleFunc("/transfers", mc.Transfers)
	mux.HandleFunc("/status", mc.GetStatus)
	mux.HandleFunc("/gossip", gossip.HandleGossip)
	http.ListenAndServe(":" + myPort, mux)
}
