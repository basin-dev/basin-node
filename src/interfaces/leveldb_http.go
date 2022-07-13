package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

func authenticate(w http.ResponseWriter, r *http.Request, db *leveldb.DB) bool {
	// u, err := db.Get([]byte("node/username"), nil)
	// if err != nil {
	// 	w.WriteHeader(500)
	// 	return false
	// }

	// p, err := db.Get([]byte("node/pass"), nil)
	// if err != nil {
	// 	w.WriteHeader(500)
	// 	return false
	// }

	// if !(bytes.Equal(u, []byte(r.URL.Query().Get("username"))) && bytes.Equal(p, []byte(r.URL.Query().Get("pass")))) {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Invalid credentials"))
	// 	return false
	// }

	// w.WriteHeader(200)

	// Going to just let everything get past for now. Assuming that the node calling the interface is trusted.
	return true
}

func handleAuth(pattern string, db *leveldb.DB, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		authenticated := authenticate(w, r, db)
		if !authenticated {
			return
		}

		handler(w, r)
	})
}

func RunHTTPServer(db *leveldb.DB) {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")

	println("Listening at address", listener.Addr())
	if err != nil {
		// Panic is when you want to share stack track trace with the programmer.
		// log.Fatal is for end user error messages.
		log.Fatal(err)
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello. The server hears your request and responds.")
	})

	handleAuth("/write", db, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(400)
			w.Write([]byte("Invalid method, " + r.Method))
		}

		type WriteBody struct {
			Url   string
			Value string
		}

		var bodyjson WriteBody
		err := json.NewDecoder(r.Body).Decode(&bodyjson)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// Then further marshal and validate value based on the schema

		err = db.Put([]byte(bodyjson.Url), []byte(bodyjson.Value), nil)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Write(true)
		w.WriteHeader(200)
	})

	handleAuth("/read", db, func(w http.ResponseWriter, r *http.Request) {

		val, err := db.Get([]byte(r.URL.Query().Get("url")), nil)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Write(val)
		w.WriteHeader(200)
	})

	log.Fatal(http.Serve(listener, nil))
}

func main() {
	db := startDB()
	RunHTTPServer(db)
}
