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
	"log"

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
		log.Println("Failed to read resource: ", err.Error())
		return Response(400, err.Error()), err
	}
	return Response(200, val), nil
}

func typecastPermissions(original []PermissionJson) []client.PermissionJson {
	// TODO: Ugh...why do I have to typecast here : (
	var permissions []client.PermissionJson
	for _, permission := range original {
		var capabilities []client.CapabilitySchema
		for _, capability := range permission.Capabilities {
			capabilities = append(capabilities, client.CapabilitySchema{Action: &capability.Action, Expiration: &capability.Expiration})
		}
		permissions = append(permissions, client.PermissionJson{Data: permission.Data, Entities: permission.Entities, Capabilities: capabilities})
	}
	return permissions
}

// Register - Register Basin resource
func (s *DefaultApiService) Register(ctx context.Context, registerRequest RegisterRequest) (ImplResponse, error) {
	permissions := typecastPermissions(registerRequest.Permissions)
	err := node.TheBasinNode.Register(ctx, registerRequest.Url, client.AdapterJson(registerRequest.Adapter), permissions, registerRequest.Schema)
	if err != nil {
		return Response(400, false), nil
	}
	return Response(200, true), nil
}

// Subscribe - Request subscription to Basin resource
func (s *DefaultApiService) Subscribe(ctx context.Context, subscribeRequest SubscribeRequest) (ImplResponse, error) {
	permissions := typecastPermissions(subscribeRequest.Permissions)
	err := node.TheBasinNode.RequestSubscriptions(ctx, permissions)
	if err != nil {
		return Response(400, err.Error()), nil
	}

	return Response(200, map[string]interface{}{}), nil
}

// Write - Write Basin resource
func (s *DefaultApiService) Write(ctx context.Context, writeRequest WriteRequest) (ImplResponse, error) {
	err := node.TheBasinNode.WriteResource(ctx, writeRequest.Url, []byte(writeRequest.Value))
	if err != nil {
		return Response(400, false), nil
	}
	return Response(200, true), nil
}
