package util

import (
	. "github.com/sestinj/basin-node/adapters"
	. "github.com/sestinj/basin-node/structs"
)

func GetPermissions(dataUrl string) *[]PermissionJson {
	url := GetMetadataUrl(dataUrl, "permissions")
	data := LocalAdapter.Read(url)
	perms := Unmarshal[[]PermissionJson](data)

	return perms
}
