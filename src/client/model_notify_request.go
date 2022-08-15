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

// NotifyRequest struct for NotifyRequest
type NotifyRequest struct {
	Url string `json:"url"`
}

// NewNotifyRequest instantiates a new NotifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyRequest(url string) *NotifyRequest {
	this := NotifyRequest{}
	this.Url = url
	return &this
}

// NewNotifyRequestWithDefaults instantiates a new NotifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyRequestWithDefaults() *NotifyRequest {
	this := NotifyRequest{}
	return &this
}

// GetUrl returns the Url field value
func (o *NotifyRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotifyRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotifyRequest) SetUrl(v string) {
	o.Url = v
}

func (o NotifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableNotifyRequest struct {
	value *NotifyRequest
	isSet bool
}

func (v NullableNotifyRequest) Get() *NotifyRequest {
	return v.value
}

func (v *NullableNotifyRequest) Set(val *NotifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyRequest(val *NotifyRequest) *NullableNotifyRequest {
	return &NullableNotifyRequest{value: val, isSet: true}
}

func (v NullableNotifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


