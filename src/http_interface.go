package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	. "github.com/sestinj/basin-node/util"
)

/*
Presumably anyone can call the node to attempt write, but will just be rejected if they don't have the right permissions.

But if we end up wanting to do some higher level authentication then this will be here.
*/
func authenticate(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func handleAuth(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		authenticated := authenticate(w, r)
		if !authenticated {
			return
		}

		handler(w, r)
	})
}

type writeBody struct {
	Url   string
	Value []byte
}

func RunHttpServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
	}

	println("Listening at address", listener.Addr())

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello. The server hears your request and responds.")
	})

	handleAuth("/requestSubscription", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.WriteHeader(500)
			w.Write([]byte("Not yet implemented"))
		default:
			w.WriteHeader(400)
			w.Write([]byte("Invalid method, " + r.Method))
		}
	})

	// TODO: Need to perform authentication on the caller. Should have their DID and a credential to prove it!!!!
	handleAuth("/write", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(400)
			w.Write([]byte("Invalid method, " + r.Method))
		}

		var bodyjson writeBody
		err := json.NewDecoder(r.Body).Decode(&bodyjson)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// TODO: Then further marshal and validate value based on the schema
		err = WriteResource(bodyjson.Url, bodyjson.Value)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error writing to resource: " + err.Error()))
			return
		}

		w.Write([]byte("true"))
		w.WriteHeader(200)
	})

	handleAuth("/read", func(w http.ResponseWriter, r *http.Request) {
		val, err := GetResource(r.URL.Query().Get("url"))
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error reading resource: " + err.Error()))
		}

		w.Write(val)
		w.WriteHeader(200)
	})

	log.Fatal(http.Serve(listener, nil))
}
