package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

type LevelDbHttpAdapter struct{}

type writeBody struct {
	Url   string
	Value []byte
}

func (l LevelDbHttpAdapter) Read(url string) []byte {
	log.Fatal("THERE A PROBLEM HERE")
	// TODO: Here lies the problem: url is supposed to be a Basin URL, but can't use that for any HTTP request. Need to get the HTTP url from metadata, but this itself requires making some request to a Basin URL...
	// So we need a way of resolving Basin URLs. I think that basically what has to happen is we resolve to the machine that stores the data based only off of the user part of the URL, and then get the data from there. Gonna require some DHT stuff.

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return resBody
}

func (l LevelDbHttpAdapter) Write(url string, value []byte) bool {
	log.Fatal("THERE A PROBLEM HERE")

	body := writeBody{Url: url, Value: value}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := new(bool)
	json.Unmarshal(resBody, result)

	return *result
}

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

func RunHttpServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
	}

	println("Listening at address", listener.Addr())

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello. The server hears your request and responds.")
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

		LocalAdapter.Write(bodyjson.Url, []byte(bodyjson.Value))

		w.Write([]byte("true"))
		w.WriteHeader(200)
	})

	handleAuth("/read", func(w http.ResponseWriter, r *http.Request) {
		val := LocalAdapter.Read(r.URL.Query().Get("url"))

		w.Write(val)
		w.WriteHeader(200)
	})

	log.Fatal(http.Serve(listener, nil))
}
