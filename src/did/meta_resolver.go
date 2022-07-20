package did

import (
	"encoding/json"
	"errors"

	. "github.com/ockam-network/did"

	goDid "github.com/nuts-foundation/go-did/did"
)

type Resolver interface {
	Resolve(did string) ([]byte, error)
}

type MetaResolver struct {
	Resolvers map[string]Resolver
}

func (m MetaResolver) Resolve(did string) (*goDid.Document, error) {
	parsed, err := Parse(did)
	if err != nil {
		return nil, err
	}

	resolver, prs := m.Resolvers[parsed.Method]
	if prs {
		rawDoc, err := resolver.Resolve(did)
		if err != nil {
			return nil, err
		}

		doc := new(goDid.Document)
		err = json.Unmarshal(rawDoc, doc)

		return doc, err
	} else {
		return nil, errors.New("Resolver for did not found")
	}
}

func NewMetaResolver(resolvers map[string]Resolver) *MetaResolver {
	return &MetaResolver{Resolvers: resolvers}
}
