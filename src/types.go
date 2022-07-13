// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "fmt"
import "reflect"
import "encoding/json"

type CapabilitySchemaJson struct {
	// Action corresponds to the JSON schema field "action".
	Action *CapabilitySchemaJsonAction `json:"action,omitempty"`

	// Expiration corresponds to the JSON schema field "expiration".
	Expiration *string `json:"expiration,omitempty"`
}

type CapabilitySchemaJsonAction string

const CapabilitySchemaJsonActionNotify CapabilitySchemaJsonAction = "notify"
const CapabilitySchemaJsonActionRead CapabilitySchemaJsonAction = "read"
const CapabilitySchemaJsonActionWrite CapabilitySchemaJsonAction = "write"

type PermissionSchemaJson struct {
	// Capabilities corresponds to the JSON schema field "capabilities".
	Capabilities []CapabilitySchemaJson `json:"capabilities,omitempty"`

	// Data corresponds to the JSON schema field "data".
	Data []string `json:"data,omitempty"`

	// Entities corresponds to the JSON schema field "entities".
	Entities []string `json:"entities,omitempty"`
}

type UrlSchemaJson struct {
	// Domain corresponds to the JSON schema field "domain".
	Domain *string `json:"domain,omitempty"`

	// Scheme corresponds to the JSON schema field "scheme".
	Scheme *string `json:"scheme,omitempty"`

	// User corresponds to the JSON schema field "user".
	User *string `json:"user,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CapabilitySchemaJsonAction) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CapabilitySchemaJsonAction {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CapabilitySchemaJsonAction, v)
	}
	*j = CapabilitySchemaJsonAction(v)
	return nil
}

var enumValues_CapabilitySchemaJsonAction = []interface{}{
	"read",
	"write",
	"notify",
}