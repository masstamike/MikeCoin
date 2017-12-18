package main

import (
	"net/http"
	"blockchain/api"
	"blockchain/api/gossip"
	"os"
	"strconv"
	"fmt"
	"time"
)

func main() {

	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Incorrect number of args, expects 1 or 2.")
		os.Exit(1)
	}
	myPort, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var peerPort int
	if len(os.Args) == 3 {
		peerPort, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		peerPort = 0
	}

	var mc api.MikeCoin
	mc.Initialize()
	gossip.Initialize(myPort, peerPort)


	mux := http.DefaultServeMux
	mux.HandleFunc("/balance", mc.Balance)
	mux.HandleFunc("/users", mc.Users)
	mux.HandleFunc("/transfers", mc.Transfers)
	mux.HandleFunc("/status", mc.GetStatus)
	mux.HandleFunc("/gossip", gossip.HandleGossip)
	go http.ListenAndServe(":" + strconv.Itoa(myPort), mux)
	for {
		time.Sleep(5000 * time.Millisecond)
		gossip.ChangeMovie()
		gossip.Gossip()
		gossip.PrintPeers()
	}
}
