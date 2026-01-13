package did

import (
	"strings"

	"io"
	"net/http"

	. "github.com/ockam-network/did"
)

type WebResolver struct{}

// As based on https://w3c-ccg.github.io/did-method-web/#read-resolve
func (w WebResolver) Resolve(did string) ([]byte, error) {
	parsed, err := Parse(did)
	if err != nil {
		return nil, err
	}

	path := strings.ReplaceAll(parsed.ID, ":", "/")
	// "If the domain contains a port percent decode the colon."

	url := "https://" + path

	// If there is not path specified, use the well-known path
	if !strings.Contains(path, "/") {
		url += "/.well-known"
	}

	url += "/did.json"

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	return resBody, err
}

func NewWebResolver() *WebResolver {
	return &WebResolver{}
}
