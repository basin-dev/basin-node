/* The metaadaptere just takes a URL and makes the correct call to the adapter that lives there.
Basically this is only responsible for deciding whether the URL is Basin, Local, HTTP, or other and then routing to the correct adapter.
The meta adapter lives on the node.
It is only a NICE TO HAVE feature right now.
*/
package adapters

import (
	"encoding/json"
	"errors"

	. "github.com/sestinj/basin-node/structs"
	"github.com/sestinj/basin-node/util"
)

type Adapter interface {
	Read(url string) ([]byte, error)
	Write(url string, value []byte) error
}

type RawAdapterConfig struct {
	AdapterName string
	Config      []byte
}

type MetaAdapter struct{}

func getAdapterConfig(dataUrl string) (RawAdapterConfig, error) {
	url := util.GetMetadataUrl(dataUrl, util.Adapter)
	cfg := new(AdapterConfig)
	var raw RawAdapterConfig
	bytes, err := LocalAdapter.Read(url)
	if err != nil {
		return raw, err
	}
	err = json.Unmarshal(bytes, cfg)
	if err != nil {
		return raw, err
	}

	cfgData, err := json.Marshal(cfg.Config)
	if err != nil {
		return raw, err
	}
	raw.AdapterName = cfg.AdapterName
	raw.Config = cfgData

	return raw, nil
}

func selectAdapter(url string) (Adapter, error) {
	cfg, err := getAdapterConfig(url)
	if err != nil {
		return nil, err
	}
	switch cfg.AdapterName {
	case "local":
		return LocalAdapter, nil
	case "http":
		return httpAdapter, nil
	default:
		// TODO: Adapter plugins
		return nil, errors.New("Unknown adapter")
	}
}

func (m MetaAdapter) Read(url string) ([]byte, error) {
	adapter, err := selectAdapter(url)
	if err != nil {
		return nil, err
	}
	return adapter.Read(url)
}

func (m MetaAdapter) Write(url string, value []byte) error {
	adapter, err := selectAdapter(url)
	if err != nil {
		return err
	}
	return adapter.Write(url, value)
}

var MainAdapter = MetaAdapter{}
