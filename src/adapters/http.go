package adapters

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type HttpAdapter struct{}

func (l HttpAdapter) Read(url string) ([]byte, error) {
	log.Fatal("THERE A PROBLEM HERE")
	// TODO: Here lies the problem: url is supposed to be a Basin URL, but can't use that for any HTTP request. Need to get the HTTP url from metadata, but this itself requires making some request to a Basin URL...
	// So we need a way of resolving Basin URLs. I think that basically what has to happen is we resolve to the machine that stores the data based only off of the user part of the URL, and then get the data from there. Gonna require some DHT stuff.

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resBody, nil
}

func (l HttpAdapter) Write(url string, value []byte) error {
	log.Fatal("THERE A PROBLEM HERE")

	body := writeBody{Url: url, Value: value}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		return err
	}
	reader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
