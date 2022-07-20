package util

import (
	"encoding/json"
	"log"
	"strings"

	. "github.com/sestinj/basin-node/structs"
)

func Unmarshal[T any](data []byte) *T {
	result := new(T)
	json.Unmarshal(data, result)

	return result
}

type ErrorHandleAction int64

/*
Panic is when you want to share stack track trace with the programmer.

log.Fatal is for end user error messages.
*/
const (
	Panic ErrorHandleAction = iota
	LogFatal
)

func HandleErr(err error, action ErrorHandleAction) {
	if err != nil {
		switch action {
		case Panic:
			panic(err)
		case LogFatal:
			log.Fatal(err)
		}
	}
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
	return "basin://" + did + ".basin." + dataName.String()
}
