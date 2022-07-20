/** All of the structs are going in one place because eventually I want to autogenerate many of them. */

package structs

import (
	"time"
)

type UrlJson struct {
	Scheme string
	User   string
	Domain string
}

type CapabilityActionOption int64

const (
	Read CapabilityActionOption = iota
	Write
)

type CapabilityJson struct {
	Action     CapabilityActionOption
	Expiration time.Time
}

type PermissionJson struct {
	Data         []string
	Capabilities []CapabilityJson
	Entities     []string
}
