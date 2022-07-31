package did

import (
	"errors"

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
		return nil, errors.New("Resolver for did not found")
	}
}

func NewMetaResolver(resolvers map[string]Resolver) *MetaResolver {
	return &MetaResolver{Resolvers: resolvers}
}
