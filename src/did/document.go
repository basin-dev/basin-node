package did

import (
	"encoding/json"

	. "github.com/ockam-network/did"

	"net/url"
)

type Controller []string

func (c *Controller) UnmarshalJSON(b []byte) error {
	if b[0] == '[' {
		return json.Unmarshal(b, (*[]string)(c))
	} else {
		var str string
		err := json.Unmarshal(b, str)
		if err != nil {
			return err
		}

		slc := new([]string)
		*slc = append(*slc, str)

		c = (*Controller)(slc)

		return err
	}
}

func (c *Controller) MarshalJSON() ([]byte, error) {
	if len(*c) == 0 {
		return nil, nil
	} else if len(*c) == 1 {
		return json.Marshal(([]string)(*c)[0])
	} else {
		return json.Marshal(([]string)(*c))
	}
}

type DidDocument struct {
	Id          DID        `json:"id"`
	Context     []url.URL  `json:"@context"`
	AlsoKnownAs []string   `json:"alsoKnownAs,omitempty"`
	Controller  Controller `json:"controller,omitempty"`
}
