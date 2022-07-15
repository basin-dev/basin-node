package util

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/xeipuuv/gojsonschema"
)

// psuedocode to just to figure out what I need to do
func getSources(mode, db) {
	switch mode {
	case "consumer":
		// query db for consumer Basin URLs
		return "consumer-urls"
	case "producer":
		// query db for producer Basin URLs
		return "producer-urls"
	default:
		return "other-urls"
}