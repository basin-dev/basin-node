package adapters

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var (
	httpAdapter = HttpAdapter{}
)

type HttpAdapter struct{}

type EndpointDescription struct {
	Url    string
	Method string
	Body   []byte
}

type HttpAdapterConfig struct {
	Read  EndpointDescription
	Write EndpointDescription
}

func parseHttpConfig(url string) (HttpAdapterConfig, error) {
	fullCfg, err := getAdapterConfig(url)
	cfg := new(HttpAdapterConfig)
	if err != nil {
		return *cfg, err
	}
	err = json.Unmarshal(fullCfg.Config, cfg)
	if err != nil {
		return *cfg, err
	}
	return *cfg, nil
}

func (l HttpAdapter) Read(url string) ([]byte, error) {
	cfg, err := parseHttpConfig(url)
	if err != nil {
		return nil, err
	}

	return performRequest(cfg.Read)
}

func (l HttpAdapter) Write(url string, value []byte) error {
	cfg, err := parseHttpConfig(url)
	if err != nil {
		return err
	}

	_, err = performRequest(cfg.Write)
	return err
}

func (l HttpAdapter) Modify(url string, value []byte) error {
	cfg, err := parseHttpConfig(url)
	if err != nil {
		return err
	}

	_, err = performRequest(cfg.Modify)
	return err
}

func performRequest(endpoint EndpointDescription) ([]byte, error) {
	reader := bytes.NewReader(endpoint.Body)

	req, err := http.NewRequest(endpoint.Method, endpoint.Url, reader)
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
