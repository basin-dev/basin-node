/* The metaadaptere just takes a URL and makes the correct call to the adapter that lives there.
Basically this is only responsible for deciding whether the URL is Basin, Local, HTTP, or other and then routing to the correct adapter.
The meta adapter lives on the node.
It is only a NICE TO HAVE feature right now.
*/
package adapters

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

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

var LOCAL_ADAPTER_CONFIG = RawAdapterConfig{AdapterName: "local"}

type MetaAdapter struct{}

func getAdapterConfig(dataUrl string) (RawAdapterConfig, error) {
	var raw RawAdapterConfig
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
		log.Printf("Unknown meta prefix in URL '%s'", parsed.Domain)
		return raw, errors.New("Unknown meta prefix")
	}

	// TODO: For now, bottoming out with user data (basin.producer....) files, but want to probably register them in the same way instead. See below
	if strings.HasPrefix(parsed.Domain, "basin") {
		return LOCAL_ADAPTER_CONFIG, nil
	}

	// TODO: TODO: TODO: When a file is written, should it's adapter info always be written (unless it's an adapter file?? This is getting ugly fast...)
	// TODO: Question you need to answer rn is whether there should exist a meta.adapter.basin.producer.sources file from the start, or if you should assume that basin.producer.sources is automatically local. What files should be local? Wouldn't we want to register this like anything else?

	url := util.GetMetadataUrl(dataUrl, util.Adapter)
	cfg := new(AdapterConfig)
	bytes, err := LocalAdapter.Read(url)
	if err != nil {
		log.Println("Error reading from local LevelDB: " + err.Error())
		return raw, err
	}
	err = json.Unmarshal(bytes, cfg)
	if err != nil {
		log.Println("Error unmarshaling config: " + err.Error())
		return raw, err
	}

	cfgData, err := json.Marshal(cfg.Config)
	if err != nil {
		log.Println("Error marshaling config: " + err.Error())
		return raw, err
	}
	raw.AdapterName = cfg.AdapterName
	raw.Config = cfgData

	return raw, nil
}

func selectAdapter(url string) (Adapter, error) {
	cfg, err := getAdapterConfig(url)
	if err != nil {
		log.Println("Error getting adapter config: " + err.Error())
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
		log.Printf("Error selecting adapter: %s\n", err.Error())
		return err
	}
	return adapter.Write(url, value)
}

var MainAdapter = MetaAdapter{}
