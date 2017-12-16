package api

import (
	"fmt"
	"net/http"
	"sync"
	"strconv"
)

type MikeCoin struct {
	balances map[string]int
}

var instance *MikeCoin
var once sync.Once

func (MikeCoin) Initialize() (*MikeCoin) {
	once.Do(func() {
		instance = &MikeCoin{}
		instance.balances = make(map[string]int)
		instance.balances["michael"] = 100000
	})
	return instance
}

func (mc *MikeCoin) PrintState () {
	fmt.Println(mc.balances)
}

func (MikeCoin) Balance (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		r.ParseForm()
		s := ""
		user := r.Form.Get("user")

		currency, ok := instance.balances[user]
		if !ok {
			http.Error(w, "Not Found", 404)
		} else {
			s += "User: " + user + ", Value: " + strconv.Itoa(currency) + "\n"
			fmt.Fprintf(w, s)
		}
	default:
		http.Error(w, "Invalid Request Method", 400)
	}
}

func (MikeCoin) Users (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		username := r.PostForm.Get("user")
		_, ok := instance.balances[username]
		if !ok {
			instance.balances[username] = 0
		}
	default:
		http.Error(w, "Invalid Request Method", 400)
	}
}

func (MikeCoin) Transfers (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		to := r.PostForm.Get("to")
		from := r.PostForm.Get("from")
		amount,err := strconv.Atoi(r.PostForm.Get("amount"))
		if err != nil {
			http.Error(w, "Invalid Amount", 400)
			return
		}

		fromBalance, ok := instance.balances[from]
		if !ok {
			http.Error(w, "From User Does Not Exist", 400)
			return
		}

		if fromBalance < amount {
			http.Error(w, "Insufficient Funds", 400)
			return
		}
		_, ok = instance.balances[to]
		if !ok {
			http.Error(w, "To User Does Not Exist", 400)
			return
		}

		instance.balances[from] -= amount
		instance.balances[to] += amount

	default:
		http.Error(w, "Invalid Request Method", 400)
	}
}

func (MikeCoin) GetStatus (w http.ResponseWriter, r *http.Request) {
	s := ""
	for key,val :=  range instance.balances {
		s += "Key: " + key + ", Value: " + strconv.Itoa(val) + "\n"
	}
	fmt.Fprintf(w, s)
}