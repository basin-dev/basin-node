/* The metaadaptere just takes a URL and makes the correct call to the adapter that lives there.
Basically this is only responsible for deciding whether the URL is Basin, Local, HTTP, or other and then routing to the correct adapter.
The meta adapter lives on the node.
It is only a NICE TO HAVE feature right now.
*/
package adapters

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sestinj/basin-node/client"
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

var LOCAL_ADAPTER_CONFIG = client.AdapterJson{AdapterName: "local"}

type MetaAdapter struct{}

func getAdapterConfig(dataUrl string) (client.AdapterJson, error) {
	var adapterCfg client.AdapterJson
	// NOTE: This becomes a problem once you start trying to work with metadata: if you want to write to meta.adapter.... you first need to read meta.adapter.meta.adapter... and so on, an infinite loop.
	// So there has to be a baseline default for metadata adapters, which makes sense.
	parsed := util.ParseUrl(dataUrl)
	if strings.HasPrefix(parsed.Domain, "meta."+util.Adapter.String()) {
		return LOCAL_ADAPTER_CONFIG, nil
	} else if strings.HasPrefix(parsed.Domain, "meta."+util.Schema.String()) {
		return LOCAL_ADAPTER_CONFIG, nil
	} else if strings.HasPrefix(parsed.Domain, "meta."+util.Permissions.String()) {
		return LOCAL_ADAPTER_CONFIG, nil
	} else if strings.HasPrefix(parsed.Domain, "meta.") {
		return adapterCfg, fmt.Errorf("Unknown meta prefix in URL '%s'", parsed.Domain)
	}

	// TODO: For now, bottoming out with user data (basin.producer....) files, but want to probably register them in the same way instead. See below
	if strings.HasPrefix(parsed.Domain, "basin") {
		return LOCAL_ADAPTER_CONFIG, nil
	}

	// TODO: TODO: TODO: When a file is written, should it's adapter info always be written (unless it's an adapter file?? This is getting ugly fast...)
	// TODO: Question you need to answer rn is whether there should exist a meta.adapter.basin.producer.sources file from the start, or if you should assume that basin.producer.sources is automatically local. What files should be local? Wouldn't we want to register this like anything else?

	url := util.GetMetadataUrl(dataUrl, util.Adapter)
	bytes, err := LocalAdapter.Read(url)
	if err != nil {
		return adapterCfg, fmt.Errorf("Error reading from local LevelDB: %w\n", err)
	}
	err = json.Unmarshal(bytes, &adapterCfg)
	if err != nil {
		return adapterCfg, fmt.Errorf("Error unmarshaling config: %w\n", err)
	}

	return adapterCfg, nil
}

func selectAdapter(url string) (Adapter, error) {
	cfg, err := getAdapterConfig(url)
	if err != nil {
		return nil, fmt.Errorf("Error getting adapter config for url %s: %w\n", url, err)
	}
	switch cfg.AdapterName {
	case "local":
		return LocalAdapter, nil
	case "http":
		return httpAdapter, nil
	default:
		// TODO: Adapter plugins
		return nil, fmt.Errorf("Unknown adapter name '%s'", cfg.AdapterName)
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
		return fmt.Errorf("Error selecting adapter for url %s: %w\n", url, err)
	}
	return adapter.Write(url, value)
}

var MainAdapter = MetaAdapter{}
