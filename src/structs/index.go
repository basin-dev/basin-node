/** All of the structs are going in one place because eventually I want to autogenerate many of them. */

package structs

type UrlJson struct {
	Scheme string `json:"scheme"`
	User   string `json:"user"`
	Domain string `json:"domain"`
}

type CapabilityJson struct {
	Action     string `json:"action"`
	Expiration string `json:"expiration"`
}

type PermissionJson struct {
	Data         []string         `json:"data"`
	Capabilities []CapabilityJson `json:"capabilities"`
	Entities     []string         `json:"entities"`
}

// TODO: Is this a totally open-ended JSON Schema document?
type SchemaJson struct {
}

type WalletInfoJson struct {
	Did string `json:"did"`
}

type ManifestJson struct {
	Url          string      `json:"url"`
	ProducerDid  string      `json:"producerDid"`
	Dependencies []string    `json:"dependencies"`
	Name         string      `json:"name"`
	Version      string      `json:"version"`
	Description  string      `json:"description"`
	Schema       interface{} `json:"interface"`
	PublicRead   bool        `json:"publicRead"`
	PublicWrite  bool        `json:"publicWrite"`
}
