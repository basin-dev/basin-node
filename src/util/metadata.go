package util

import (
	. "github.com/sestinj/basin-node/adapters"
	. "github.com/sestinj/basin-node/structs"
)

func GetWalletInfo() *WalletInfoJson {
	data := LocalAdapter.Read("local://wallet")

	return Unmarshal[WalletInfoJson](data)
}

func GetPermissions(dataUrl string) *[]PermissionJson {
	url := GetMetadataUrl(dataUrl, "permissions")
	mdata := LocalAdapter.Read(url)

	return Unmarshal[[]PermissionJson](mdata)
}

func GetSchema(dataUrl string) *SchemaJson {
	url := GetMetadataUrl(dataUrl, "schema")

	mdata := LocalAdapter.Read(url)
	return Unmarshal[SchemaJson](mdata)
}

func GetSources(mode string) *[]string {
	walletInfo := GetWalletInfo()

	url := GetUserDataUrl(walletInfo.Did, mode+".urls")
	mdata := LocalAdapter.Read(url)

	return Unmarshal[[]string](mdata)
}

func GetSchemas(mode string) *[]SchemaJson {
	sources := GetSources(mode)

	var schemas []SchemaJson

	for _, source := range *sources {
		schema := GetSchema(source)
		schemas = append(schemas, *schema)
	}

	return schemas
}
