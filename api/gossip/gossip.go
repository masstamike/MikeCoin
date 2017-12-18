package gossip

import (
	"net/http"
	"encoding/json"
	"strconv"
	"bytes"
	"fmt"
	"math/rand"
)


type peer struct {
	Port int `json:"port"`
	State state `json:"state"`
}

type state struct {
	Movie string `json:"movie"`
	Version int `json:"version"`
}

var myPeers []peer
var movies []string
var selfState peer

func Initialize(selfPort int, peerPort int) {
	selfState = peer{
		Port: selfPort,
		State: state{
			Movie: "Get Out",
			Version: 1,
		},
	}

	myPeers = make([]peer, 0)

	if peerPort != 0 {
		p := peer{
			Port: peerPort,
		}
		myPeers = append(myPeers, p)
	}

	movies = []string {
		"Master and Commander",
		"Get Out",
		"National Lampoon's Christmas Vacation",
		"Star Wars: The Last Jedi",
	}
}

func ChangeMovie () {
	selfState.State.Movie = movies[rand.Intn(len(movies))]
	selfState.State.Version++
}

func Gossip () {
	for p := range myPeers {
		peerUrl := "http://localhost:" + strconv.Itoa(myPeers[p].Port) +
			"/gossip"
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(append(myPeers, selfState))
		resp, err := http.Post(peerUrl, "application/json", b)
		if err != nil {
			fmt.Println(err.Error())
		}
		var otherPeers []peer
		json.NewDecoder(resp.Body).Decode(&otherPeers)
		updatePeers(otherPeers)
	}
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
			json.NewEncoder(w).Encode(append(myPeers, selfState))
		}
	default:
		http.Error(w, "Invalid Request Method", 400)
	}
}

func PrintPeers () {
	fmt.Println(append(myPeers, selfState))
}

func updatePeers (otherPeers []peer) {
	for p := range otherPeers {
		exists := false
		for myPeer := range myPeers {
			if otherPeers[p].Port == myPeers[myPeer].Port {
				exists = true
				if otherPeers[p].State.Version > myPeers[myPeer].State.Version {
					myPeers[myPeer].State = otherPeers[p].State
				}
			}
		}
		if !exists && otherPeers[p].Port != 0 && otherPeers[p].Port != selfState.Port {
			myPeers = append(myPeers, otherPeers[p])
		}
	}
}
