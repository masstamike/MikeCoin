package gossip

import (
	"net/http"
	"encoding/json"
)

type peer struct {
	Port int `json:"port"`
}

var myPeers []peer

func Initialize(peerPort int) {
	p := peer{
		Port: peerPort,
	}
	myPeers = append(make([]peer,0), p)
}

func Gossip () {

}

func HandleGossip (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var otherPeers []peer
		r.ParseForm()
		err := json.NewDecoder(r.Body).Decode(&otherPeers)
		if err != nil {
			http.Error(w, "Invalid Request Body. Please send peers.", 400)
		} else {
			updatePeers(otherPeers)
			json.NewEncoder(w).Encode(myPeers)
		}
	default:
		http.Error(w, "Invalid Request Method", 400)
	}
}

func updatePeers (otherPeers []peer) {
	for p := range otherPeers {
		exists := false
		for myPeer := range myPeers {
			if otherPeers[p] == myPeers[myPeer] {
				exists = true
			}
		}
		if !exists {
			myPeers = append(myPeers, otherPeers[p])
		}
	}
}
