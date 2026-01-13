package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/sestinj/basin-node/structs"
	"gopkg.in/yaml.v2"
)

func Unmarshal[T any](data []byte) *T {
	result := new(T)
	json.Unmarshal(data, result)

	return result
}

/* Read the contents of the file, either yaml or json, and unmarshal to the specified type */
func UnmarshalFromFile[T any](filepath string) (*T, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	t := new(T)

	segs := strings.Split(filepath, ".")
	switch segs[len(segs)-1] {
	case "json":
		err = json.Unmarshal(raw, &t)
	case "yaml":
		err = yaml.Unmarshal(raw, &t)
	case "yml":
		err = yaml.Unmarshal(raw, &t)
	default:
		return nil, fmt.Errorf("Cannot parse filetype '%s'", segs[len(segs)-1])
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}

/* Takes a file with raw bytes (in either yaml or json format), converts to the given type T, then encodes this back into bytes but in JSON format.
This allows users to input whatever markdown filetype they feel most comfortable with, but it is always converted to a standard format: JSON.
*/
func MarshalToJson[T any](filepath string) ([]byte, error) {
	t, err := UnmarshalFromFile[T](filepath)
	if err != nil {
		return nil, err
	}

	rawJson, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return rawJson, nil
}

func ParseUrl(url string) *UrlJson {

	urlJson := new(UrlJson)

	segs := strings.Split(url, "://")
	urlJson.Scheme = segs[0]

	segs = strings.Split(segs[1], ".")
	urlJson.User = segs[0]
	urlJson.Domain = strings.Join(segs[1:], ".")

	return urlJson
}

func PrintUrl(url *UrlJson) string {
	return url.Scheme + "://" + url.User + "." + url.Domain
}

func GetMetadataUrl(dataUrl string, prefix MetadataPrefix) string {
	parsed := ParseUrl(dataUrl)
	parsed.Domain = "meta." + prefix.String() + "." + parsed.Domain
	return PrintUrl(parsed)
}

func GetUserDataUrl(did string, dataName string) string {
	return "basin://" + did + ".basin." + dataName
}

// basin://com.natesesti.com.fb.firstname
// basin://com.natesesti.meta.schema.fb.firstname
// basin://com.natesesti.basin.producers.sources

// basin://com.twitter.users.tydunn

func Contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}
