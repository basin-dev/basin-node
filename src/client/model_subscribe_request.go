/*
Basin RPC API

Basin RPC API

API version: 1.0.11
Contact: sestinj@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubscribeRequest struct for SubscribeRequest
type SubscribeRequest struct {
	Permissions []PermissionJson `json:"permissions"`
}

// NewSubscribeRequest instantiates a new SubscribeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeRequest(permissions []PermissionJson) *SubscribeRequest {
	this := SubscribeRequest{}
	this.Permissions = permissions
	return &this
}

// NewSubscribeRequestWithDefaults instantiates a new SubscribeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeRequestWithDefaults() *SubscribeRequest {
	this := SubscribeRequest{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *SubscribeRequest) GetPermissions() []PermissionJson {
	if o == nil {
		var ret []PermissionJson
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *SubscribeRequest) GetPermissionsOk() ([]PermissionJson, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *SubscribeRequest) SetPermissions(v []PermissionJson) {
	o.Permissions = v
}

func (o SubscribeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableSubscribeRequest struct {
	value *SubscribeRequest
	isSet bool
}

func (v NullableSubscribeRequest) Get() *SubscribeRequest {
	return v.value
}

func (v *NullableSubscribeRequest) Set(val *SubscribeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeRequest(val *SubscribeRequest) *NullableSubscribeRequest {
	return &NullableSubscribeRequest{value: val, isSet: true}
}

func (v NullableSubscribeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


