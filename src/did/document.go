/*
The contents of this file are no longer being used, but might be useful later if we decide to make our own library.
Now using github.com/nuts-foundation/go-did/did
*/

package did

import (
	"encoding/json"

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

// TODO[LIBRARY][1]: See https://www.w3.org/TR/did-core/#dfn-publickeyjwk, https://www.rfc-editor.org/rfc/rfc7517
type Jwk interface{}

type VerificationMethod struct {
	Id                 string `json:"id"`
	Controller         string `json:"controller"`
	Type               string `json:"type"`
	PublicKeyJwk       Jwk    `json:"publicKeyJwk,omitempty"`
	PublicKeyMultibase string `json:"publicKeyMultibase,omitempty"`
}

// https://www.w3.org/TR/did-core/#did-document-properties
type DidDocument struct {
	Id                   string               `json:"id"` // Is there a way to automatically parse this with a custom unmarshaler? Can't add method to non-local struct, but maybe wrapper
	Context              []url.URL            `json:"@context"`
	AlsoKnownAs          []string             `json:"alsoKnownAs,omitempty"`
	Controller           Controller           `json:"controller,omitempty"`
	VerificationMethod   []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication       []string             `json:"authentication,omitempty"`
	AssertionMethod      []string             `json:"assertionMethod,omitempty"`
	CapabilityDelegation []string             `json:"capabilityDelegation,omitempty"`
	CapabilityInvocation []string             `json:"capabilityInvocation,omitempty"`
	KeyAgreement         []VerificationMethod `json:"keyAgreement,omitempty"`
}
