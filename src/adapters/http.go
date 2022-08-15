package adapters

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
	log "github.com/sestinj/basin-node/log"
)

var (
	httpAdapter = HttpAdapter{}
)

type HttpAdapter struct{}

type EndpointDescription struct {
	Url    string `json:"url"`
	Method string `json:"method"`
	Body   string `json:"body"`
}

type HttpAdapterConfig struct {
	Read  EndpointDescription `json:"read"`
	Write EndpointDescription `json:"write"`
}

func parseHttpConfig(url string) (HttpAdapterConfig, error) {
	fullCfg, err := getAdapterConfig(url)
	log.Info.Println(fullCfg)
	cfg := new(HttpAdapterConfig)
	if err != nil {
		return *cfg, err
	}
	err = mapstructure.Decode(fullCfg.Config, cfg)
	if err != nil {
		return *cfg, fmt.Errorf("Error decoding config: %w", err)
	}
	return *cfg, nil
}

func (l HttpAdapter) Read(url string) ([]byte, error) {
	cfg, err := parseHttpConfig(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse HTTP config for url %s: %w", url, err)
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

func performRequest(endpoint EndpointDescription) ([]byte, error) {
	reader := strings.NewReader(endpoint.Body)

	req, err := http.NewRequest(endpoint.Method, endpoint.Url, reader)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
