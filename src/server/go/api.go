/*
 * Basin RPC API
 *
 * Basin RPC API
 *
 * API version: 1.0.11
 * Contact: sestinj@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"context"
	"net/http"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	Read(http.ResponseWriter, *http.Request)
	Register(http.ResponseWriter, *http.Request)
	Subscribe(http.ResponseWriter, *http.Request)
	Write(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	Read(context.Context, string) (ImplResponse, error)
	Register(context.Context, RegisterRequest) (ImplResponse, error)
	Subscribe(context.Context, SubscribeRequest) (ImplResponse, error)
	Write(context.Context, WriteRequest) (ImplResponse, error)
}
