// This file contains functions for CLI commands that work with the manifest.json file.

package scratch

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aymerick/raymond"
)

func loadManifest(path string) ManifestJson {
	manifestRaw, err := ioutil.ReadFile(path)
	if err != nil {
		handleErr(err, LogFatal)
	}

	manifest := new(ManifestJson)
	err = json.Unmarshal(manifestRaw, manifest)
	if err != nil {
		handleErr(err, LogFatal)
	}
	return manifest
}

func publishManifest(manifestPath string, resourceUrl string) {
	url := getMetadataUrl(resourceUrl, "manifest")
	permissionsUrl := getMetadataUrl(resourceUrl, "permissions")

	// Set any default values and set permissions key/value

	// Just like any other update - should tell subscribers (want a function for this)
}

type Dependency struct {
	Url          string
	SchemaName   string
	ManifestName string
}

func generateListener(manifestPath string, outputPath string) {
	manifest := loadManifest(manifestPath)

	// First, generate the types, using gojsonschema???

	// Then, generate a stub for each type - what does the developer do when new data comes in? This is just the first step. They probably don't want to generate their codebase, then they'll either want to 1) set off a lambda function, call some specified URL with an HTTP request, or maybe have a special IDE for writing the functions
	// How do I run a shell command from go code?

	template, err := ioutil.ReadFile("../templates/listener.go.mustache")
	if err != nil {
		handleErr(err, LogFatal)
	}

	ctx := []Dependency{
		{
			Url:          "https://",
			SchemaName:   "abcd",
			ManifestName: "abcde",
		},
	}

	result, err := raymond.Render(string(template), ctx)
	handleErr(err, LogFatal)

	err = ioutil.WriteFile(outputPath, []byte(result), 0)
	handleErr(err, LogFatal)

	for _, dep := range manifest.dependencies {
		// Download the manifest of the dependency to get info?
		// Remove spaces in the name, and disambiguate (replace spaces with _)

		append(ctx, Dependency{})
	}

	// The listener code can be the same for everyone

	// But you have to generate the code that unmarshals data, validates type, and sends it strongly typed to the stubbed functions
}

// Writing data to a domain should depend on both the collector (domain owner) and user (first part of domain). The domain owner should be able to specify that they store the data but that the user can sign the message when they want to write

// Script to generate a strongly typed object from a URL, or from a manifest
// Start listener from manifest
// Royalty payment protocol with ethereum
// Some work on royalty schemes—how are they described? And the algorithm to decide payment amounts
// A package manager website - but this also serves as your homepage to control your own data, manage unions, etc...
