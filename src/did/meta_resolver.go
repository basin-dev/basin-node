package did

import (
	"fmt"

	. "github.com/ockam-network/did"
)

type Resolver interface {
	/* Takes a DID string, resolves with the given method, and returns []byte which can be unmarshaled to the DID document */
	Resolve(did string) ([]byte, error)
}

type MetaResolver struct {
	Resolvers map[string]Resolver
}

func (m MetaResolver) Resolve(did string) ([]byte, error) {
	parsed, err := Parse(did)
	if err != nil {
		return nil, err
	}

	resolver, prs := m.Resolvers[parsed.Method]
	if prs {
		doc, err := resolver.Resolve(did)
		if err != nil {
			return nil, err
		}

		return doc, err
	} else {
		return nil, fmt.Errorf("Resolver for DID %s not found", did)
	}
}

func NewMetaResolver(resolvers map[string]Resolver) *MetaResolver {
	return &MetaResolver{Resolvers: resolvers}
}
