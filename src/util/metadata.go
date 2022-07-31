package util

type MetadataPrefix int64

const (
	Permissions MetadataPrefix = iota
	Royalties
	Schema
	Manifest
)

func (m MetadataPrefix) String() string {
	switch m {
	case Permissions:
		return "permissions"
	case Royalties:
		return "royalties"
	case Schema:
		return "schema"
	case Manifest:
		return "manifest"
	default:
		return "unknown"
	}
}
