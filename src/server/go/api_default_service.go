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
	"errors"
	"net/http"

	"github.com/sestinj/basin-node/client"
	"github.com/sestinj/basin-node/node"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// Read - Read Basin resource
func (s *DefaultApiService) Read(ctx context.Context, url string) (ImplResponse, error) {
	val, err := node.TheBasinNode.ReadResource(ctx, url)
	if err != nil {
		return Response(400, nil), nil
	}
	return Response(200, val), nil
}

// Register - Register Basin resource
func (s *DefaultApiService) Register(ctx context.Context, registerRequest RegisterRequest) (ImplResponse, error) {
	// TODO: Ugh...why do I have to typecast here : (
	var permissions []client.PermissionJson
	for _, permission := range registerRequest.Permissions {
		var capabilities []client.CapabilitySchema
		for _, capability := range permission.Capabilities {
			capabilities = append(capabilities, client.CapabilitySchema{Action: &capability.Action, Expiration: &capability.Expiration})
		}
		permissions = append(permissions, client.PermissionJson{Data: permission.Data, Entities: permission.Entities, Capabilities: capabilities})
	}
	err := node.TheBasinNode.Register(ctx, registerRequest.Url, client.AdapterJson(registerRequest.Adapter), permissions, registerRequest.Schema)
	if err != nil {
		return Response(400, false), nil
	}
	return Response(200, true), nil
}

// Subscribe - Request subscription to Basin resource
func (s *DefaultApiService) Subscribe(ctx context.Context, subscribeRequest SubscribeRequest) (ImplResponse, error) {
	// err := node.TheBasinNode.RequestSubscription(ctx, subscribeRequest.Url, subscribeRequest.Permissions)
	// TODO: Weird thing you decided to do...
	return Response(http.StatusNotImplemented, nil), errors.New("Subscribe method not implemented")
}

// Write - Write Basin resource
func (s *DefaultApiService) Write(ctx context.Context, writeRequest WriteRequest) (ImplResponse, error) {
	err := node.TheBasinNode.WriteResource(ctx, writeRequest.Url, []byte(writeRequest.Value))
	if err != nil {
		return Response(400, false), nil
	}
	return Response(200, true), nil
}
